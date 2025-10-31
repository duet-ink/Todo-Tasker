# Todo-Tasker Development Roadmap

Version-wise todos for building a production-ready application.

**Current Version:** v0.0.7

---

## Tasks

### Backend

#### Foundation
- [x] Fix Makefile `deploy.yml` reference (should be `compose.yml`)
- [x] Add proper MIME types for WASM files
- [ ] Add basic request logging middleware
- [ ] Test WASM file serving in browser
- [ ] Verify SQLite WASM initialization
- [ ] Add CORS headers if needed for WASM loading

#### WebSocket & Sync
- [ ] Add WebSocket upgrade capability (Gorilla WebSocket)
- [ ] Implement WebSocket sync handler (server/sync.go)
- [ ] Create in-memory session storage (no persistence)
- [ ] Implement JOIN message handling
- [ ] Implement METADATA_EXCHANGE handling
- [ ] Implement REQUEST_SYNC handling
- [ ] Implement SYNC_DATA handling
- [ ] Implement SYNC_COMPLETE handling
- [ ] Implement message relay between devices
- [ ] Add session cleanup (remove after 15 minutes)
- [ ] Add keep-alive PING/PONG mechanism
- [ ] Add WebSocket route to server
- [ ] Test WebSocket connection and message routing
- [ ] Test multi-device sync (3+ devices)
- [ ] Add connection timeout handling
- [ ] Implement rate limiting for sync messages
- [ ] Optimize WebSocket message routing
- [ ] Add compression for large sync payloads
- [ ] Implement WebSocket connection pooling

#### Security
- [ ] Implement security middleware
  - [ ] Content Security Policy (CSP) headers
  - [ ] X-Frame-Options
  - [ ] X-Content-Type-Options
  - [ ] HSTS headers
- [ ] Add rate limiting middleware
  - [ ] Limit sync sessions per IP
  - [ ] Limit WebSocket messages
- [ ] Add input validation middleware
- [ ] Implement request size limits
- [ ] Add panic recovery middleware
- [ ] Test with malicious inputs

#### Monitoring & Health
- [ ] Implement health check endpoint (/health)
- [ ] Implement readiness endpoint (/ready)
- [ ] Add metrics endpoint (/metrics)
- [ ] Add Prometheus metrics
  - [ ] Active sync sessions
  - [ ] Messages relayed
  - [ ] Average session duration
  - [ ] Error rates
- [ ] Improve error logging with context
- [ ] Add structured logging throughout
- [ ] Configure log aggregation
- [ ] Add logging for sync analytics

#### Performance
- [ ] Implement message batching
- [ ] Add gzip compression middleware
- [ ] Optimize memory usage
- [ ] Add connection limits
- [ ] Implement HTTP/2 server push
- [ ] Profile and optimize hot paths
- [ ] Add caching for static assets

#### Testing
- [ ] Write unit tests for sync logic
  - [ ] Session management
  - [ ] Message routing
  - [ ] Session cleanup
- [ ] Write integration tests
  - [ ] WebSocket connection
  - [ ] Full sync flow
  - [ ] Multi-device scenarios
- [ ] Add test coverage reporting (target: 80%+)
- [ ] Add Go linting (golangci-lint)
- [ ] Fix all linter warnings
- [ ] Add pre-commit hooks

#### Documentation
- [ ] Add godoc comments to all public functions
- [ ] Document sync protocol in code
- [ ] Create API documentation
- [ ] Add architecture diagrams

#### Deployment
- [ ] Create multi-stage Docker build
- [ ] Minimize Docker image size
- [ ] Add health checks to containers
- [ ] Implement graceful shutdown
- [ ] Add container security scanning
- [ ] Configure production logging
- [ ] Set up log rotation
- [ ] Add service worker route
- [ ] Configure cache headers for static assets

#### CI/CD
- [ ] Expand GitHub Actions workflows
  - [ ] Run tests on PR
  - [ ] Build Docker images
  - [ ] Security scanning
  - [ ] Linting checks
- [ ] Add automated releases
- [ ] Implement deployment status checks
- [ ] Add rollback mechanism
- [ ] Configure deployment notifications

### Frontend/UI/UX

#### Foundation
- [ ] Complete todo.html component structure
- [ ] Complete dashboard.html component structure
- [ ] Complete settings.html component structure
- [ ] Fix responsive behavior on mobile devices
- [ ] Add loading states for HTMX requests
- [ ] Test all components load via HTMX

#### SQLite WASM Setup
- [ ] Initialize SQLite WASM on page load
- [ ] Create database schema
  - [ ] todos table (id, title, description, completed, priority, due_date, device_id, created_at, updated_at, deleted, synced)
  - [ ] device_info table (device_id, device_name, created_at, last_sync)
  - [ ] sync_log table (session_id, remote_device_id, direction, todos_sent, todos_received, status, started_at, completed_at, error_message)
- [ ] Implement IndexedDB persistence for SQLite
- [ ] Add database migration system
- [ ] Create device ID generation and storage (localStorage)
- [ ] Add database error handling and user feedback
- [ ] Test database operations in browser console
- [ ] Test WASM loading in different browsers

#### Todo CRUD Operations
- [ ] Implement "Add Todo" functionality
  - [ ] Input field with validation
  - [ ] Generate composite ID (deviceId-timestamp-random)
  - [ ] Insert into SQLite database
  - [ ] Update UI immediately
- [ ] Implement "Display Todos" functionality
  - [ ] Query all non-deleted todos
  - [ ] Render todo list with Alpine.js
  - [ ] Add empty state message
- [ ] Implement "Toggle Complete" functionality
  - [ ] Update completed status
  - [ ] Update updated_at timestamp
  - [ ] Refresh UI
- [ ] Implement "Delete Todo" functionality
  - [ ] Soft delete (set deleted=1)
  - [ ] Update updated_at timestamp
  - [ ] Remove from UI
- [ ] Add todo editing functionality
- [ ] Test CRUD operations work offline

#### Todo Features
- [ ] Add todo filtering (all, active, completed)
- [ ] Add todo sorting options (date, priority, status)
- [ ] Implement search functionality
- [ ] Add todo priority levels (low, medium, high)
- [ ] Add due dates with date picker
- [ ] Add todo categories/tags
- [ ] Implement bulk operations
  - [ ] Select multiple todos
  - [ ] Bulk delete
  - [ ] Bulk complete
- [ ] Add undo functionality
- [ ] Add keyboard shortcuts

#### Sync Implementation
- [ ] Create SyncEngine class (sync-engine.js)
- [ ] Implement QR code generation
  - [ ] Add qrcode.js library
  - [ ] Generate sync metadata (topId, lastModified, todoCount)
  - [ ] Create session ID
  - [ ] Display QR code on screen
- [ ] Implement QR code scanner
  - [ ] Add html5-qrcode library
  - [ ] Camera access and QR scanning UI
  - [ ] Parse QR data and validate
- [ ] Implement WebSocket client connection
- [ ] Implement metadata exchange logic
  - [ ] Get all todo IDs from database
  - [ ] Send metadata to remote device
  - [ ] Receive and compare remote metadata
- [ ] Implement sync request logic
  - [ ] Identify missing todos
  - [ ] Request specific todo IDs
- [ ] Implement sync data handling
  - [ ] Receive todo data from remote
  - [ ] Merge todos into local database
  - [ ] Handle conflict resolution (Last-Write-Wins)
- [ ] Implement conflict resolution
  - [ ] Compare updated_at timestamps
  - [ ] Update local todo if remote is newer
  - [ ] Keep local todo if it's newer
- [ ] Handle deleted todos sync
  - [ ] Send deleted todos to remote
  - [ ] Process remote deletions locally
- [ ] Implement sync log
  - [ ] Record sync sessions in database
  - [ ] Track todos sent/received
  - [ ] Display sync history
- [ ] Add retry logic for failed syncs

#### Sync UI
- [ ] Add sync UI components
  - [ ] "Sync with another device" button
  - [ ] QR code display modal
  - [ ] QR scanner modal
  - [ ] Sync status indicator
- [ ] Add sync progress UI
  - [ ] Show "Syncing X of Y todos"
  - [ ] Progress bar
  - [ ] Success/failure messages
- [ ] Add session timeout handling
- [ ] Test sync with various scenarios
  - [ ] Both devices have new todos
  - [ ] One device has deletions
  - [ ] Conflicting modifications
  - [ ] Large dataset (100+ todos)

#### Dashboard
- [ ] Display todo statistics
  - [ ] Total todos
  - [ ] Completed todos
  - [ ] Pending todos
  - [ ] Completion rate
- [ ] Show sync status and history
- [ ] Display recent activity log
- [ ] Add charts/visualizations

#### Settings
- [ ] Device name configuration
- [ ] Theme selection (light/dark mode)
- [ ] Data export functionality (JSON, CSV)
- [ ] Data import functionality
- [ ] Clear all data option
- [ ] View sync history
- [ ] Display storage usage

#### PWA & Offline
- [ ] Complete manifest.json
  - [ ] Add app icons (192x192, 512x512)
  - [ ] Configure start_url, display mode
  - [ ] Set theme colors
  - [ ] Add shortcuts
- [ ] Implement service worker
  - [ ] Cache static assets (CSS, JS, WASM)
  - [ ] Cache HTML pages and components
  - [ ] Add offline fallback page
  - [ ] Implement cache versioning
- [ ] Add install prompt for PWA
- [ ] Add offline indicator in UI
- [ ] Test offline functionality
  - [ ] App loads when offline
  - [ ] CRUD operations work offline
  - [ ] UI shows offline status
- [ ] Test PWA installation on mobile devices
- [ ] Run Lighthouse PWA audit (target: 90+)

#### Security & Validation
- [ ] Sanitize all user inputs (XSS prevention)
- [ ] Validate todo data before insertion
- [ ] Add client-side error boundary
- [ ] Implement graceful error handling
  - [ ] Database errors
  - [ ] Network errors
  - [ ] Sync errors
- [ ] Add user-friendly error messages
- [ ] Implement error recovery mechanisms

#### Performance
- [ ] Implement virtual scrolling for large lists
- [ ] Add debouncing for search/filter
- [ ] Optimize SQLite queries (add indexes)
- [ ] Minimize JavaScript bundle size
- [ ] Implement lazy loading for components
- [ ] Add image optimization
- [ ] Optimize database vacuum scheduling
- [ ] Add batch sync for large datasets
- [ ] Run Lighthouse performance audit (target: 90+)
- [ ] Measure and optimize Time to Interactive (TTI)
- [ ] Performance testing
  - [ ] Test with 1000+ todos
  - [ ] Measure sync time
  - [ ] Check memory usage
  - [ ] Fix memory leaks

#### Testing
- [ ] Write JavaScript unit tests
  - [ ] SyncEngine class
  - [ ] Conflict resolution logic
  - [ ] Database operations
  - [ ] ID generation
- [ ] Write integration tests
  - [ ] Todo CRUD operations
  - [ ] Sync workflow
  - [ ] Offline mode
- [ ] Add E2E tests (Playwright/Cypress)
  - [ ] Create and complete todos
  - [ ] Sync between devices
  - [ ] PWA installation
- [ ] Add visual regression tests
- [ ] Cross-browser testing
  - [ ] Chrome
  - [ ] Firefox
  - [ ] Safari
  - [ ] Edge
- [ ] Mobile device testing
  - [ ] iOS Safari
  - [ ] Android Chrome

#### Styling & UX
- [ ] Add basic styling with DaisyUI
- [ ] UI/UX consistency review
- [ ] Add micro-interactions
- [ ] Polish animations and transitions
- [ ] Add loading skeletons
- [ ] Complete empty states for all views
- [ ] Improve mobile responsiveness
- [ ] Final accessibility audit (WCAG 2.1 AA)
  - [ ] Keyboard navigation
  - [ ] ARIA labels
  - [ ] Screen reader testing
  - [ ] Color contrast
  - [ ] Focus indicators

#### Documentation
- [ ] Document SyncEngine API
- [ ] Create user guide with screenshots
- [ ] Write FAQ section
- [ ] Create troubleshooting guide
- [ ] Add keyboard shortcuts reference
- [ ] Create video tutorial
- [ ] Document database schema
- [ ] Add JSDoc comments
- [ ] Create component documentation

#### Analytics & Monitoring
- [ ] Add client-side analytics (privacy-respecting)
  - [ ] Feature usage tracking
  - [ ] Sync success rates
  - [ ] Error rates
  - [ ] Performance metrics
- [ ] Configure error tracking (Sentry/similar)
- [ ] Add user feedback mechanism

#### Production Preparation
- [ ] Add production build script
- [ ] Minify assets
- [ ] Add source maps for debugging
- [ ] Configure CDN for static assets
- [ ] Add feature flags system
- [ ] Add frontend build to CI
- [ ] Add JavaScript linting to CI
- [ ] Add test coverage checks
- [ ] Add bundle size monitoring
- [ ] Configure deployment previews

#### Polish
- [ ] Add onboarding tour for new users
- [ ] Create sample data generator
- [ ] Add command palette (Cmd+K)
- [ ] Polish all error messages
- [ ] Final mobile optimization
- [ ] Add app tour walkthrough

---

*Last Updated: 2025-10-31*
