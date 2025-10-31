# Sync Architecture - Offline-First Decentralized Todo Sync

**Design Goal:** Build an offline-first, decentralized sync system where the server stores **zero user data** permanently.

**Last Updated:** 2025-10-31

---

## Table of Contents

- [Overview](#overview)
- [Core Principles](#core-principles)
- [Sync Protocol](#sync-protocol)
- [Technical Architecture](#technical-architecture)
- [Data Structures](#data-structures)
- [Implementation Details](#implementation-details)
- [Conflict Resolution](#conflict-resolution)
- [Security Considerations](#security-considerations)
- [Edge Cases](#edge-cases)

---

## Overview

The Todo-Tasker sync system uses a **QR code handshake + WebSocket data transfer** approach where:

1. Device A displays a QR code with sync metadata
2. Device B scans the QR code and establishes WebSocket connection
3. Devices exchange metadata to determine who has newer data
4. Missing/updated todos are transferred directly between devices
5. Server acts only as a **message relay** (signaling server)
6. **No user data is stored on the server** - it only routes messages

### Key Benefits

- ✅ **Zero server storage costs** - Server never persists user data
- ✅ **Privacy-first** - Data flows directly between user's devices
- ✅ **Offline-first** - Each device has complete local copy
- ✅ **Cost-efficient** - Minimal server bandwidth usage
- ✅ **Scalable** - Server just routes messages, no database queries
- ✅ **Simple** - No complex authentication or session management

---

## Core Principles

### 1. Eventual Consistency
Devices eventually have the same data, but may temporarily diverge when offline.

### 2. Client-Side Storage
All data stored in SQLite WASM (IndexedDB) in the browser.

### 3. Device Identity
Each browser/device has a unique `deviceId` (UUID stored in localStorage).

### 4. Composite IDs
Todo IDs use format: `{deviceId}-{timestamp}-{random}` to prevent collisions.

### 5. Last-Write-Wins (LWW)
Conflicts resolved using `updated_at` timestamp - newest wins.

### 6. Soft Deletes
Deleted todos marked with `deleted=1` flag, synced to other devices.

### 7. Server as Dumb Relay
Server only routes WebSocket messages, never reads or stores content.

---

## Sync Protocol

### Step-by-Step Flow

```
┌─────────────┐                                    ┌─────────────┐
│  Device A   │                                    │  Device B   │
│  (Primary)  │                                    │ (Secondary) │
└──────┬──────┘                                    └──────┬──────┘
       │                                                  │
       │ ① Generate QR Code                              │
       │    {sessionId, deviceId, metadata}              │
       │    Display on screen                            │
       │                                                  │
       │                                    ② Scan QR ◄──┤
       │                                       Parse data │
       │                                                  │
       │                     ③ WebSocket Connect         │
       │◄─────────────────────────────────────────────────┤
       │                    ws://server/sync              │
       │                                                  │
       │ ④ JOIN session (via sessionId)                  │
       │◄─────────────────────────────────────────────────┤
       │                                                  │
       │ ⑤ Exchange Metadata                             │
       ├─────────METADATA_EXCHANGE─────────────────────►│
       │  {topId, lastModified, todoCount}               │
       │◄─────────METADATA_EXCHANGE─────────────────────┤
       │                                                  │
       │ ⑥ Determine Sync Direction                      │
       │    (who's ahead, what's missing)                │
       │                                                  │
       │ ⑦ Request Missing Data                          │
       ├─────────REQUEST_SYNC──────────────────────────►│
       │  {needIds: [...]}                               │
       │◄─────────REQUEST_SYNC──────────────────────────┤
       │                                                  │
       │ ⑧ Send Todo Data                                │
       ├─────────SYNC_DATA─────────────────────────────►│
       │  [{id, title, completed, ...}, ...]             │
       │◄─────────SYNC_DATA─────────────────────────────┤
       │                                                  │
       │ ⑨ Merge & Resolve Conflicts                     │
       │    (Last-Write-Wins based on updated_at)        │
       │                                                  │
       │ ⑩ Acknowledge Completion                        │
       ├─────────SYNC_COMPLETE─────────────────────────►│
       │◄─────────SYNC_COMPLETE─────────────────────────┤
       │                                                  │
       │ ⑪ Close WebSocket                               │
       └──────────────────────────────────────────────────┘

Server: Only routes messages between Device A ↔ Device B
        Never reads, stores, or modifies the data
```

---

## Technical Architecture

### System Components

```
┌─────────────────────────────────────────────────────────┐
│                    Browser (Device A)                    │
│                                                           │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │   UI Layer   │  │  Sync Engine │  │  QR Generator│  │
│  │  (Alpine.js) │  │ (JavaScript) │  │  (qrcode.js) │  │
│  └──────┬───────┘  └───────┬──────┘  └──────────────┘  │
│         │                   │                             │
│  ┌──────┴───────────────────┴──────┐                    │
│  │   SQLite WASM Database          │                    │
│  │   (IndexedDB Persistence)       │                    │
│  └──────────────┬──────────────────┘                    │
│                 │                                         │
│         ┌───────┴────────┐                               │
│         │  WebSocket     │                               │
│         │  Client        │                               │
│         └───────┬────────┘                               │
└─────────────────┼───────────────────────────────────────┘
                  │
                  │ ws://server/sync
                  │
┌─────────────────┼───────────────────────────────────────┐
│                 │    Go WebSocket Server                 │
│         ┌───────┴────────┐                               │
│         │  Sync Handler  │                               │
│         │  (In-Memory    │                               │
│         │   Sessions)    │                               │
│         └───────┬────────┘                               │
│                 │                                         │
│    ┌────────────┴────────────┐                          │
│    │  Message Router         │                          │
│    │  (sessionId -> conns)   │                          │
│    └─────────────────────────┘                          │
│                                                           │
│    No Database | No Persistence | Pure Relay             │
└───────────────────────────────────────────────────────────┘
                  │
                  │ ws://server/sync
                  │
┌─────────────────┼───────────────────────────────────────┐
│                 │    Browser (Device B)                  │
│         ┌───────┴────────┐                               │
│         │  WebSocket     │                               │
│         │  Client        │                               │
│         └───────┬────────┘                               │
│                 │                                         │
│  ┌──────────────┴──────────────────┐                    │
│  │   SQLite WASM Database          │                    │
│  │   (IndexedDB Persistence)       │                    │
│  └──────┬───────────────────┬──────┘                    │
│         │                   │                             │
│  ┌──────┴───────┐  ┌───────┴──────┐  ┌──────────────┐  │
│  │   UI Layer   │  │  Sync Engine │  │  QR Scanner  │  │
│  │  (Alpine.js) │  │ (JavaScript) │  │  (html5-qr)  │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
└───────────────────────────────────────────────────────────┘
```

---

## Data Structures

### QR Code Payload

```javascript
{
  // Session identifiers
  sessionId: "uuid-v4",           // Unique sync session ID
  deviceId: "device-uuid",        // Originating device ID

  // Sync metadata
  topId: "device-a-1698765432-001", // Highest todo ID
  lastModified: 1698765432000,      // Unix timestamp (ms)
  todoCount: 42,                    // Total todos (non-deleted)

  // Connection info
  wsUrl: "ws://localhost:8080/sync", // WebSocket URL
  expires: 1698765732000             // QR expires in 5 minutes
}
```

### Todo Schema (SQLite)

```sql
CREATE TABLE todos (
  -- Primary key (composite format: deviceId-timestamp-random)
  id TEXT PRIMARY KEY,

  -- Todo data
  title TEXT NOT NULL,
  description TEXT,
  completed INTEGER DEFAULT 0,      -- 0 = false, 1 = true
  priority TEXT DEFAULT 'medium',   -- low, medium, high
  due_date INTEGER,                 -- Unix timestamp or NULL

  -- Sync metadata
  device_id TEXT NOT NULL,          -- Which device created it
  created_at INTEGER NOT NULL,      -- Unix timestamp (ms)
  updated_at INTEGER NOT NULL,      -- For conflict resolution
  deleted INTEGER DEFAULT 0,        -- Soft delete flag
  synced INTEGER DEFAULT 0          -- Has been synced at least once
);

-- Index for sync queries
CREATE INDEX idx_updated_at ON todos(updated_at);
CREATE INDEX idx_device_id ON todos(device_id);
CREATE INDEX idx_deleted ON todos(deleted);
```

### Device Schema (SQLite)

```sql
CREATE TABLE device_info (
  device_id TEXT PRIMARY KEY,       -- UUID for this device
  device_name TEXT,                 -- User-friendly name
  created_at INTEGER NOT NULL,      -- When device first used app
  last_sync INTEGER                 -- Last successful sync timestamp
);
```

### Sync Log Schema (SQLite)

```sql
CREATE TABLE sync_log (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  session_id TEXT NOT NULL,         -- Sync session ID
  remote_device_id TEXT NOT NULL,   -- Device synced with
  direction TEXT NOT NULL,          -- 'send' or 'receive'
  todos_sent INTEGER DEFAULT 0,     -- Number of todos sent
  todos_received INTEGER DEFAULT 0, -- Number of todos received
  status TEXT NOT NULL,             -- 'success', 'failed', 'partial'
  started_at INTEGER NOT NULL,      -- Sync start time
  completed_at INTEGER,             -- Sync completion time
  error_message TEXT                -- Error details if failed
);

CREATE INDEX idx_sync_session ON sync_log(session_id);
```

### WebSocket Message Format

```javascript
// Base message structure
{
  type: "MESSAGE_TYPE",      // Message type (see below)
  sessionId: "uuid",         // Sync session ID
  fromDevice: "device-uuid", // Sender device ID
  toDevice: "device-uuid",   // Recipient device ID (optional)
  timestamp: 1698765432000,  // Message timestamp
  payload: {}                // Type-specific payload
}

// Message Types:

// 1. JOIN - Join a sync session
{
  type: "JOIN",
  sessionId: "session-uuid",
  fromDevice: "device-b-uuid",
  payload: {
    deviceName: "Device B"
  }
}

// 2. METADATA_EXCHANGE - Exchange sync metadata
{
  type: "METADATA_EXCHANGE",
  sessionId: "session-uuid",
  fromDevice: "device-a-uuid",
  payload: {
    topId: "device-a-1698765432-001",
    lastModified: 1698765432000,
    todoCount: 42,
    todoIds: ["id1", "id2", "id3", ...]  // All todo IDs
  }
}

// 3. REQUEST_SYNC - Request specific todos
{
  type: "REQUEST_SYNC",
  sessionId: "session-uuid",
  fromDevice: "device-b-uuid",
  payload: {
    needIds: ["id1", "id5", "id10"]  // Todos I don't have
  }
}

// 4. SYNC_DATA - Send todo data
{
  type: "SYNC_DATA",
  sessionId: "session-uuid",
  fromDevice: "device-a-uuid",
  payload: {
    todos: [
      {
        id: "device-a-1698765432-001",
        title: "Buy groceries",
        description: "Milk, eggs, bread",
        completed: 0,
        priority: "high",
        due_date: null,
        device_id: "device-a-uuid",
        created_at: 1698765432000,
        updated_at: 1698765432000,
        deleted: 0
      },
      // ... more todos
    ]
  }
}

// 5. SYNC_COMPLETE - Acknowledge sync completion
{
  type: "SYNC_COMPLETE",
  sessionId: "session-uuid",
  fromDevice: "device-b-uuid",
  payload: {
    todosReceived: 15,
    todosSent: 8,
    status: "success"
  }
}

// 6. ERROR - Error occurred
{
  type: "ERROR",
  sessionId: "session-uuid",
  fromDevice: "device-a-uuid",
  payload: {
    code: "SYNC_FAILED",
    message: "Connection timeout"
  }
}

// 7. PING/PONG - Keep connection alive
{
  type: "PING",
  sessionId: "session-uuid",
  timestamp: 1698765432000
}
```

---

## Implementation Details

### Frontend: Sync Engine (JavaScript)

#### Generate QR Code

```javascript
// sync-engine.js

import QRCode from 'qrcode';
import { v4 as uuidv4 } from 'uuid';

class SyncEngine {
  constructor(db) {
    this.db = db;  // SQLite WASM instance
    this.ws = null;
    this.sessionId = null;
    this.deviceId = this.getOrCreateDeviceId();
  }

  getOrCreateDeviceId() {
    let deviceId = localStorage.getItem('deviceId');
    if (!deviceId) {
      deviceId = uuidv4();
      localStorage.setItem('deviceId', deviceId);
    }
    return deviceId;
  }

  async generateSyncQR() {
    // Get sync metadata from database
    const metadata = await this.getSyncMetadata();

    // Create sync session
    this.sessionId = uuidv4();

    const qrData = {
      sessionId: this.sessionId,
      deviceId: this.deviceId,
      topId: metadata.topId,
      lastModified: metadata.lastModified,
      todoCount: metadata.todoCount,
      wsUrl: `ws://${window.location.host}/sync`,
      expires: Date.now() + (5 * 60 * 1000)  // 5 minutes
    };

    // Generate QR code
    const qrCodeDataUrl = await QRCode.toDataURL(JSON.stringify(qrData));

    // Start WebSocket server and wait for scan
    await this.startSyncSession();

    return qrCodeDataUrl;
  }

  async getSyncMetadata() {
    const result = await this.db.exec(`
      SELECT
        MAX(id) as topId,
        MAX(updated_at) as lastModified,
        COUNT(*) as todoCount
      FROM todos
      WHERE deleted = 0
    `);

    return {
      topId: result[0]?.topId || null,
      lastModified: result[0]?.lastModified || 0,
      todoCount: result[0]?.todoCount || 0
    };
  }

  async getAllTodoIds() {
    const result = await this.db.exec(`
      SELECT id FROM todos ORDER BY id
    `);
    return result.map(row => row.id);
  }
}
```

#### Scan QR and Sync

```javascript
// sync-engine.js (continued)

class SyncEngine {
  async scanAndSync(qrDataString) {
    try {
      const qrData = JSON.parse(qrDataString);

      // Validate QR data
      if (Date.now() > qrData.expires) {
        throw new Error('QR code expired');
      }

      // Connect to WebSocket
      this.sessionId = qrData.sessionId;
      await this.connectWebSocket(qrData.wsUrl);

      // Join sync session
      await this.sendMessage({
        type: 'JOIN',
        payload: {
          deviceName: this.getDeviceName()
        }
      });

      // Exchange metadata
      const myMetadata = await this.getSyncMetadata();
      const myTodoIds = await this.getAllTodoIds();

      await this.sendMessage({
        type: 'METADATA_EXCHANGE',
        payload: {
          topId: myMetadata.topId,
          lastModified: myMetadata.lastModified,
          todoCount: myMetadata.todoCount,
          todoIds: myTodoIds
        }
      });

      // Wait for remote metadata and start sync
      // (handled in WebSocket message handler)

    } catch (error) {
      console.error('Sync failed:', error);
      throw error;
    }
  }

  async connectWebSocket(wsUrl) {
    return new Promise((resolve, reject) => {
      this.ws = new WebSocket(wsUrl);

      this.ws.onopen = () => {
        console.log('WebSocket connected');
        resolve();
      };

      this.ws.onerror = (error) => {
        console.error('WebSocket error:', error);
        reject(error);
      };

      this.ws.onmessage = (event) => {
        this.handleMessage(JSON.parse(event.data));
      };

      this.ws.onclose = () => {
        console.log('WebSocket closed');
      };
    });
  }

  async handleMessage(message) {
    switch (message.type) {
      case 'METADATA_EXCHANGE':
        await this.handleMetadataExchange(message.payload);
        break;

      case 'REQUEST_SYNC':
        await this.handleSyncRequest(message.payload);
        break;

      case 'SYNC_DATA':
        await this.handleSyncData(message.payload);
        break;

      case 'SYNC_COMPLETE':
        await this.handleSyncComplete(message.payload);
        break;

      case 'ERROR':
        console.error('Sync error:', message.payload);
        break;

      case 'PONG':
        // Keep-alive response
        break;
    }
  }

  async handleMetadataExchange(remoteMetadata) {
    const myTodoIds = await this.getAllTodoIds();
    const remoteTodoIds = remoteMetadata.todoIds;

    // Find missing todos
    const needFromRemote = remoteTodoIds.filter(id => !myTodoIds.includes(id));
    const remoteNeeds = myTodoIds.filter(id => !remoteTodoIds.includes(id));

    // Request missing todos
    if (needFromRemote.length > 0) {
      await this.sendMessage({
        type: 'REQUEST_SYNC',
        payload: { needIds: needFromRemote }
      });
    }

    // Send todos that remote needs
    if (remoteNeeds.length > 0) {
      const todosToSend = await this.getTodosByIds(remoteNeeds);
      await this.sendMessage({
        type: 'SYNC_DATA',
        payload: { todos: todosToSend }
      });
    }

    // If nothing to sync, complete
    if (needFromRemote.length === 0 && remoteNeeds.length === 0) {
      await this.completeSyncSession();
    }
  }

  async handleSyncRequest(payload) {
    const { needIds } = payload;
    const todos = await this.getTodosByIds(needIds);

    await this.sendMessage({
      type: 'SYNC_DATA',
      payload: { todos }
    });
  }

  async handleSyncData(payload) {
    const { todos } = payload;

    for (const todo of todos) {
      await this.mergeTodo(todo);
    }

    // Log sync
    console.log(`Received ${todos.length} todos`);

    // Check if sync complete
    await this.checkSyncCompletion();
  }

  async mergeTodo(remoteTodo) {
    // Check if todo exists locally
    const localTodo = await this.getTodoById(remoteTodo.id);

    if (!localTodo) {
      // New todo, insert it
      await this.insertTodo(remoteTodo);
    } else {
      // Conflict resolution: Last-Write-Wins
      if (remoteTodo.updated_at > localTodo.updated_at) {
        // Remote is newer, update local
        await this.updateTodo(remoteTodo);
      }
      // else: local is newer, keep it
    }
  }

  async getTodosByIds(ids) {
    const placeholders = ids.map(() => '?').join(',');
    const result = await this.db.exec(
      `SELECT * FROM todos WHERE id IN (${placeholders})`,
      ids
    );
    return result;
  }

  async insertTodo(todo) {
    await this.db.exec(`
      INSERT INTO todos (id, title, description, completed, priority,
                         due_date, device_id, created_at, updated_at, deleted)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, [
      todo.id, todo.title, todo.description, todo.completed, todo.priority,
      todo.due_date, todo.device_id, todo.created_at, todo.updated_at, todo.deleted
    ]);
  }

  async updateTodo(todo) {
    await this.db.exec(`
      UPDATE todos
      SET title = ?, description = ?, completed = ?, priority = ?,
          due_date = ?, updated_at = ?, deleted = ?
      WHERE id = ?
    `, [
      todo.title, todo.description, todo.completed, todo.priority,
      todo.due_date, todo.updated_at, todo.deleted, todo.id
    ]);
  }

  async completeSyncSession() {
    await this.sendMessage({
      type: 'SYNC_COMPLETE',
      payload: {
        status: 'success',
        timestamp: Date.now()
      }
    });

    // Update last sync time
    await this.db.exec(`
      UPDATE device_info SET last_sync = ? WHERE device_id = ?
    `, [Date.now(), this.deviceId]);

    // Close WebSocket
    this.ws.close();
  }

  sendMessage(message) {
    const fullMessage = {
      ...message,
      sessionId: this.sessionId,
      fromDevice: this.deviceId,
      timestamp: Date.now()
    };

    this.ws.send(JSON.stringify(fullMessage));
  }
}

export default SyncEngine;
```

### Backend: WebSocket Server (Go)

```go
// server/sync.go

package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: Add proper origin check in production
	},
}

type SyncMessage struct {
	Type       string                 `json:"type"`
	SessionID  string                 `json:"sessionId"`
	FromDevice string                 `json:"fromDevice"`
	ToDevice   string                 `json:"toDevice,omitempty"`
	Timestamp  int64                  `json:"timestamp"`
	Payload    map[string]interface{} `json:"payload"`
}

type SyncSession struct {
	SessionID   string
	Connections map[string]*websocket.Conn
	CreatedAt   time.Time
	mu          sync.RWMutex
}

// In-memory session storage (no persistence!)
var (
	sessions   = make(map[string]*SyncSession)
	sessionsMu sync.RWMutex
)

// Cleanup old sessions every 10 minutes
func init() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			cleanupSessions()
		}
	}()
}

func cleanupSessions() {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	now := time.Now()
	for sessionID, session := range sessions {
		// Remove sessions older than 15 minutes
		if now.Sub(session.CreatedAt) > 15*time.Minute {
			session.mu.Lock()
			for _, conn := range session.Connections {
				conn.Close()
			}
			session.mu.Unlock()
			delete(sessions, sessionID)
			slog.Info("Cleaned up expired session", "sessionId", sessionID)
		}
	}
}

func syncWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("WebSocket upgrade failed", "error", err)
		return
	}
	defer conn.Close()

	var sessionID string
	var deviceID string

	// Handle messages
	for {
		var msg SyncMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			slog.Error("Read message failed", "error", err)
			break
		}

		sessionID = msg.SessionID
		deviceID = msg.FromDevice

		switch msg.Type {
		case "JOIN":
			handleJoin(sessionID, deviceID, conn, &msg)

		case "METADATA_EXCHANGE", "REQUEST_SYNC", "SYNC_DATA", "SYNC_COMPLETE", "ERROR":
			// Relay message to other devices in session
			relayMessage(sessionID, deviceID, &msg)

		case "PING":
			// Respond with PONG
			pong := SyncMessage{
				Type:      "PONG",
				SessionID: sessionID,
				Timestamp: time.Now().UnixMilli(),
			}
			conn.WriteJSON(pong)
		}
	}

	// Remove connection on disconnect
	removeConnection(sessionID, deviceID)
}

func handleJoin(sessionID, deviceID string, conn *websocket.Conn, msg *SyncMessage) {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	// Get or create session
	session, exists := sessions[sessionID]
	if !exists {
		session = &SyncSession{
			SessionID:   sessionID,
			Connections: make(map[string]*websocket.Conn),
			CreatedAt:   time.Now(),
		}
		sessions[sessionID] = session
		slog.Info("Created new sync session", "sessionId", sessionID)
	}

	// Add connection to session
	session.mu.Lock()
	session.Connections[deviceID] = conn
	session.mu.Unlock()

	slog.Info("Device joined session",
		"sessionId", sessionID,
		"deviceId", deviceID,
		"totalDevices", len(session.Connections))

	// Notify other devices about new join
	relayMessage(sessionID, deviceID, msg)
}

func relayMessage(sessionID, fromDevice string, msg *SyncMessage) {
	sessionsMu.RLock()
	session, exists := sessions[sessionID]
	sessionsMu.RUnlock()

	if !exists {
		slog.Warn("Session not found", "sessionId", sessionID)
		return
	}

	session.mu.RLock()
	defer session.mu.RUnlock()

	// Relay to all other devices in session
	for deviceID, conn := range session.Connections {
		if deviceID != fromDevice {
			err := conn.WriteJSON(msg)
			if err != nil {
				slog.Error("Failed to relay message",
					"error", err,
					"toDevice", deviceID)
			}
		}
	}
}

func removeConnection(sessionID, deviceID string) {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	session, exists := sessions[sessionID]
	if !exists {
		return
	}

	session.mu.Lock()
	delete(session.Connections, deviceID)
	remainingDevices := len(session.Connections)
	session.mu.Unlock()

	slog.Info("Device left session",
		"sessionId", sessionID,
		"deviceId", deviceID,
		"remainingDevices", remainingDevices)

	// Remove session if no devices left
	if remainingDevices == 0 {
		delete(sessions, sessionID)
		slog.Info("Removed empty session", "sessionId", sessionID)
	}
}
```

Add route to `server/server.go`:

```go
func New() *http.ServeMux {
	return routes{
		"/":         indexPage,
		"/sync":     syncWebSocketHandler,  // Add this
	}.getCommonRoutes().createRoutes()
}
```

---

## Conflict Resolution

### Last-Write-Wins (LWW) Strategy

When the same todo exists on both devices with different data:

```javascript
async mergeTodo(remoteTodo) {
  const localTodo = await this.getTodoById(remoteTodo.id);

  if (!localTodo) {
    // New todo, insert it
    await this.insertTodo(remoteTodo);
    return;
  }

  // Conflict detected - use updated_at timestamp
  if (remoteTodo.updated_at > localTodo.updated_at) {
    // Remote version is newer
    await this.updateTodo(remoteTodo);
    console.log(`Resolved conflict: remote version newer for ${remoteTodo.id}`);
  } else if (remoteTodo.updated_at < localTodo.updated_at) {
    // Local version is newer, keep it
    console.log(`Resolved conflict: local version newer for ${remoteTodo.id}`);
  } else {
    // Same timestamp (rare) - use device ID as tiebreaker
    if (remoteTodo.device_id < localTodo.device_id) {
      await this.updateTodo(remoteTodo);
    }
  }
}
```

### Handling Deletions

```javascript
async handleDeletedTodo(deletedTodo) {
  const localTodo = await this.getTodoById(deletedTodo.id);

  if (!localTodo) {
    // Insert tombstone record
    await this.insertTodo(deletedTodo);
  } else if (!localTodo.deleted) {
    // Local has active version, check timestamp
    if (deletedTodo.updated_at > localTodo.updated_at) {
      // Deletion is newer, mark as deleted
      await this.markTodoDeleted(deletedTodo.id, deletedTodo.updated_at);
    }
  }
}
```

---

## Security Considerations

### 1. QR Code Expiry
- QR codes expire after 5 minutes
- Prevents old QR codes from being reused
- Session IDs are single-use

### 2. Session Isolation
- Each sync session has unique `sessionId`
- Server only routes messages within same session
- No cross-session data leakage

### 3. Client-Side Encryption (Future)
- Optional: Encrypt todos before transmission
- Server can't read encrypted data
- Each user has encryption key

```javascript
// Future enhancement
async encryptTodo(todo, key) {
  const encoder = new TextEncoder();
  const data = encoder.encode(JSON.stringify(todo));
  const iv = crypto.getRandomValues(new Uint8Array(12));

  const encrypted = await crypto.subtle.encrypt(
    { name: 'AES-GCM', iv },
    key,
    data
  );

  return { encrypted, iv };
}
```

### 4. Rate Limiting
- Limit sync sessions per IP
- Limit message size and frequency
- Prevent DoS attacks

### 5. Input Validation
- Validate all message types
- Sanitize todo content
- Prevent XSS in todo titles/descriptions

---

## Edge Cases

### 1. Network Disconnection During Sync

**Problem:** WebSocket disconnects mid-sync

**Solution:**
- Mark sync as 'partial' in sync_log
- Next sync detects incomplete state
- Resume from where it left off

```javascript
async handleDisconnect() {
  await this.db.exec(`
    UPDATE sync_log
    SET status = 'partial', completed_at = ?
    WHERE session_id = ? AND status = 'in_progress'
  `, [Date.now(), this.sessionId]);
}
```

### 2. Same Todo Modified on Both Devices While Offline

**Problem:** Todo "Buy milk" changed to "Buy whole milk" on Device A, "Buy 2% milk" on Device B

**Solution:**
- Last-Write-Wins based on `updated_at`
- Newer timestamp wins
- User can manually fix if needed

### 3. Device Clock Skew

**Problem:** Device B has clock 1 hour behind, all its updates appear "older"

**Solution:**
- Use server timestamp when possible
- Validate timestamp sanity (not too far in past/future)
- Warn user about clock skew

```javascript
validateTimestamp(ts) {
  const now = Date.now();
  const oneDay = 24 * 60 * 60 * 1000;

  if (Math.abs(ts - now) > oneDay) {
    console.warn('Timestamp appears invalid, possible clock skew');
    // Use current time instead
    return now;
  }
  return ts;
}
```

### 4. Multiple Devices Syncing Simultaneously

**Problem:** Device A, B, C all try to sync at once

**Solution:**
- Server supports N devices in one session
- Messages broadcast to all devices in session
- Each device merges incoming data independently

### 5. Large Todo List (1000+ items)

**Problem:** Sending 1000 todos over WebSocket is slow

**Solution:**
- Batch sync in chunks of 100 todos
- Show progress bar
- Use compression (gzip WebSocket messages)

```javascript
async sendTodosInBatches(todos, batchSize = 100) {
  for (let i = 0; i < todos.length; i += batchSize) {
    const batch = todos.slice(i, i + batchSize);
    await this.sendMessage({
      type: 'SYNC_DATA',
      payload: {
        todos: batch,
        batchIndex: Math.floor(i / batchSize),
        totalBatches: Math.ceil(todos.length / batchSize)
      }
    });

    // Small delay between batches
    await new Promise(resolve => setTimeout(resolve, 100));
  }
}
```

### 6. Malformed QR Code

**Problem:** User scans corrupted or invalid QR code

**Solution:**
- Validate JSON parsing
- Check required fields exist
- Show user-friendly error message

```javascript
validateQRData(qrData) {
  if (!qrData.sessionId || !qrData.deviceId || !qrData.wsUrl) {
    throw new Error('Invalid QR code: missing required fields');
  }

  if (Date.now() > qrData.expires) {
    throw new Error('QR code has expired. Please generate a new one.');
  }

  return true;
}
```

---

## Future Enhancements

### 1. Continuous Sync Mode
- Keep WebSocket connection open
- Auto-sync changes in real-time
- Battery/bandwidth aware

### 2. Conflict UI
- Show user when conflicts occur
- Let user choose which version to keep
- Merge tool for manual resolution

### 3. Group Sync
- Multiple users share todo lists
- Permission system (owner, editor, viewer)
- Activity log

### 4. Backup to User's Cloud
- Optional Google Drive/Dropbox sync
- Encrypted backup files
- Restore from backup

### 5. Sync Analytics
- Track sync success rate
- Monitor sync duration
- Detect sync issues early

---

## Testing Strategy

### Unit Tests
- Test conflict resolution logic
- Test ID generation (no collisions)
- Test soft delete handling
- Test metadata calculation

### Integration Tests
- Test full sync flow between two devices
- Test network disconnect/reconnect
- Test multiple device sync
- Test large dataset sync (1000+ todos)

### Performance Tests
- Measure sync time for various dataset sizes
- Monitor WebSocket message sizes
- Check memory usage during sync
- Verify no memory leaks

### Security Tests
- Validate session isolation
- Test expired QR codes rejected
- Verify no data leakage between sessions
- Test malicious message handling

---

## Monitoring & Metrics

### Server Metrics
- Active sync sessions count
- Messages relayed per second
- Average session duration
- Session cleanup rate

### Client Metrics
- Sync success/failure rate
- Average sync duration
- Todos synced per session
- Network errors during sync

---

*End of Sync Architecture Document*
