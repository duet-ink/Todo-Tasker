# Getting Started with Todo-Tasker

Todo-Tasker is a lightweight, cost-efficient todo management application built with Go, HTMX, Alpine.js, and client-side SQLite WASM.

## Project Goal

**Build a scalable multi-user todo application with total hosting costs under $10/month.**

### Why Todo-Tasker?

Todo-Tasker is designed around a unique architecture that eliminates traditional database costs:

- **Client-Side Storage**: SQLite WASM runs entirely in the browser (IndexedDB), eliminating database server costs
- **Stateless Backend**: Go server only serves HTML and assets, enabling cheap horizontal scaling
- **Single Binary Deployment**: All assets embedded at compile time for simplified deployment
- **Minimal Dependencies**: No Redis, PostgreSQL, or external services required
- **Efficient Resources**: Lightweight Go binary can serve thousands of users on minimal hardware

**Expected Cost Breakdown**:
- VPS hosting: $5-6/month (1GB RAM sufficient)
- Domain: ~$1/month
- Database: $0 (client-side only)
- CDN: $0 (uses free public CDNs)
- APIs: $0 (no external services)

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running Locally](#running-locally)
- [Docker Deployment](#docker-deployment)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Development](#development)
- [Testing](#testing)

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.23.3 or higher** - [Download Go](https://golang.org/dl/)
- **Git** - [Download Git](https://git-scm.com/downloads)
- **Docker** (optional, for containerized deployment) - [Download Docker](https://www.docker.com/get-started)
- **Make** (optional, for using Makefile commands)
- **GCC/musl-dev** (required for SQLite CGO support)

## Installation

### Clone the Repository

```bash
git clone https://github.com/HarudaySharma/Todo-Tasker.git
cd Todo-Tasker
```

### Install Dependencies

```bash
go mod download
```

This will download all required Go dependencies.

## Running Locally

### Using Make (Recommended)

The easiest way to run the application locally is using Make:

```bash
make dev
```

This starts the application on port `8080`. Open your browser and navigate to:

```
http://localhost:8080
```

### Using Go Commands

Alternatively, you can run the application directly with Go:

```bash
go run .
```

By default, this will:
- Start the main server on port `80` (requires sudo/admin privileges)
- Start the admin server on port `4657`

### Custom Port Configuration

To run on different ports, set environment variables:

```bash
PORT=8080 ADMIN_PORT=4660 go run .
```

Or create an `app.env` file in the project root (see `.env.example`):

```env
PORT=8080
ADMIN_PORT=4660
```

## Docker Deployment

### Using Make Commands

Build and run with Docker Compose:

```bash
make run
```

This command:
- Uses `config/compose.yml` for configuration
- Builds from `config/Dockerfile`
- Exposes port `11000` internally
- Mounts volume `deploy_sqlite_dir` to `/data`
- Connects to external network `deploy_net`

To stop the containers:

```bash
make stop
```

### Manual Docker Commands

```bash
# Build the image
docker build -f config/Dockerfile -t todo-tasker .

# Run the container
docker run -p 8080:80 -p 4660:4657 todo-tasker
```

## Project Structure

```
Todo-Tasker/
├── app.go                  # Main application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies checksums
├── Makefile                # Build and development commands
├── README.md               # Project overview
│
├── config/                 # Configuration and deployment
│   ├── config.go          # Environment config and logging
│   ├── Dockerfile         # Docker build configuration
│   └── compose.yml        # Docker Compose setup
│
├── server/                 # HTTP server and routing
│   ├── server.go          # Server setup and route registration
│   ├── routes.go          # HTTP handlers
│   │
│   ├── pages/             # HTML templates
│   │   ├── layout.html    # Base layout template
│   │   ├── index.html     # Main application page
│   │   ├── admin.html     # Admin interface
│   │   ├── error.html     # Error page template
│   │   └── components/    # Reusable HTMX components
│   │       ├── assistent.html
│   │       ├── dashboard.html
│   │       ├── footer.html
│   │       ├── login.html
│   │       ├── navbar.html
│   │       ├── settings.html
│   │       ├── sidebar.html
│   │       └── todo.html
│   │
│   ├── assets/            # Static files
│   │   ├── styles.css     # Application styles
│   │   ├── manifest.json  # PWA manifest
│   │   ├── robots.txt     # Search engine directives
│   │   ├── images/        # Images and icons
│   │   │   ├── favicon.ico
│   │   │   └── logo.webp
│   │   └── database/      # SQLite WASM files
│   │       ├── sqlite3.js
│   │       ├── sqlite3.wasm
│   │       └── worker.js
│   │
│   └── middleware/        # HTTP middleware
│       ├── database.go    # Database middleware
│       ├── redis.go       # Redis middleware
│       └── security.go    # Security middleware
│
├── tests/                  # Test files
│   ├── components_test.go # Component endpoint tests
│   ├── config_test.go     # Configuration tests
│   ├── middleware_test.go # Middleware tests
│   └── server_test.go     # Server tests
│
├── docs/                   # Documentation
│   └── getting-started.md # This file
│
├── ansible/                # Deployment automation
│   └── playbook.yml       # Ansible deployment script
│
└── .github/               # CI/CD pipelines
    └── workflows/         # GitHub Actions
```

## Usage

### Main Application

Once the server is running, access the main application at:

```
http://localhost:8080/
```

The application provides:
- Todo list management with client-side SQLite storage
- Real-time updates using HTMX
- Responsive UI with DaisyUI/Tailwind CSS
- All data stored locally in your browser

### Admin Interface

Access the admin panel at:

```
http://localhost:4660/admin/
```

### Available Routes

- `/` - Main application (todo list interface)
- `/admin/` - Admin interface
- `/404/` - 404 error page
- `/error/` - Generic error page
- `POST /c/{name}` - Dynamic component loading endpoint
- `/assets/*` - Static assets (CSS, JS, images)

### Available Components

Components can be loaded dynamically via HTMX at `/c/{name}`:

- `navbar` - Navigation bar
- `sidebar` - Side navigation
- `footer` - Page footer
- `dashboard` - Main dashboard view
- `todo` - Todo list component
- `settings` - Settings panel
- `login` - Login form
- `assistent` - Assistant interface

## Development

### Available Make Commands

```bash
make dev      # Run locally on port 8080
make run      # Build and run with Docker Compose
make stop     # Stop Docker containers
make test     # Run all tests
```

### File Watching and Hot Reload

For development with automatic reloading, use [Air](https://github.com/cosmtrek/air):

```bash
go install github.com/cosmtrek/air@latest
air
```

### Making Changes

#### Adding New Routes

Edit `server/server.go` (server.go:80) and add routes to the `New()` or `NewAdmin()` functions:

```go
func New() *http.ServeMux {
    return routes{
        "/":       indexPage,
        "/new/":   newPage,  // Add your new route
        "/404/":   pageNotFound,
        "/error/": errorPage,
        "POST /c/{name}": componentsPage,
    }.createRoutes()
}
```

Then implement the handler in `server/routes.go`.

#### Adding Components

Create a new HTML file in `server/pages/components/`:

```bash
touch server/pages/components/my-component.html
```

Load it via HTMX:

```html
<div hx-post="/c/my-component" hx-trigger="load" hx-swap="outerHTML"></div>
```

The component name corresponds to the filename (without `.html` extension).

#### Modifying Templates

Templates are embedded at compile time using `//go:embed`. After modifying any HTML in `server/pages/`, restart the application:

```bash
# Ctrl+C to stop, then:
go run .
```

## Testing

### Run All Tests

```bash
make test
# or
go test ./tests/ -v
```

### Run Tests with Coverage

```bash
go test -cover ./tests/
```

### Run Specific Test File

```bash
go test -v ./tests/server_test.go
```

### Run Tests with Race Detection

```bash
go test -race ./tests/
```

### Test Organization

- `tests/config_test.go` - Configuration and environment variable tests
- `tests/server_test.go` - Server route and handler tests (routes.go:8-55)
- `tests/middleware_test.go` - Middleware functionality tests
- `tests/components_test.go` - Component endpoint tests (routes.go:8)

## Architecture Overview

### Two-Server Design

Todo-Tasker runs two concurrent HTTP servers (app.go:15):

1. **Main Server** (default port 80)
   - Serves the public-facing application
   - Handles todo list functionality
   - Routes defined in `server.New()`

2. **Admin Server** (default port 4657)
   - Administrative interface
   - Separate from main application
   - Routes defined in `server.NewAdmin()`

### Frontend Stack

All frontend libraries are loaded from CDN:

- **HTMX 2.0.7**: Dynamic HTML loading without full page refreshes
- **Alpine.js 3.x**: Client-side reactivity and interactivity
- **DaisyUI 5**: Pre-built UI components
- **Tailwind CSS 4**: Utility-first styling
- **Tabler Icons**: Icon library
- **SQLite WASM**: Client-side database (server/assets/database/)

### Template System

Templates use Go's `html/template` package:

- **Base Layout**: `server/pages/layout.html` with `{{template "main" .}}` block
- **Pages**: Define content using `{{define "main"}}...{{end}}`
- **Components**: Rendered on-demand via `/c/{name}` endpoint (routes.go:8)
- **Embedded**: All templates embedded at compile time (server.go:35-40)

### Data Flow

1. Client requests HTML page (e.g., `/`)
2. Server executes Go template with layout (routes.go:17-24)
3. Page loads with HTMX directives
4. HTMX triggers component requests to `/c/{name}` (routes.go:8)
5. Server returns rendered component HTML
6. Client-side Alpine.js and SQLite handle interactivity and data storage

### Client-Side Database Strategy

**All data storage happens in the browser:**

- SQLite WASM runs entirely client-side using IndexedDB
- WASM files located in `server/assets/database/`
- No server-side database required
- Each user's data stored locally in their browser
- Privacy benefit: No user data stored on servers
- Cost benefit: Eliminates database hosting costs

**Trade-offs:**

- Data is local to each browser/device (no automatic cross-device sync)
- Users responsible for their own data backups
- Future feature: Optional cloud sync could be added

## Key Features

### Embedded Assets

All HTML templates and static files are embedded at compile time (server.go:35-40):

```go
//go:embed pages
_pages embed.FS

//go:embed assets
_assets embed.FS
```

**Benefits:**
- Single binary contains everything
- No need to distribute separate asset files
- Faster deployment
- Simpler Docker images

**Important:** Changes to templates or assets require recompilation.

### Component Loading Pattern

Dynamic components are loaded via HTMX (routes.go:8-14):

```html
<div hx-post="/c/navbar" hx-trigger="load" hx-swap="outerHTML"></div>
```

The server's `componentsPage` handler uses the path parameter to determine which component to render.

### Cost Optimization Design

When adding features, always consider:

1. **Can this run client-side?** (Use WASM/JS instead of backend)
2. **Can we use free CDN resources?** (Don't host libraries yourself)
3. **Can we avoid external services?** (No Redis, PostgreSQL, S3, etc.)
4. **Can we reduce bandwidth?** (Minimize API calls, compress assets)
5. **Can we stay stateless?** (Enables cheap horizontal scaling)

**Anti-patterns to avoid:**
- Adding server-side database (breaks cost model)
- Adding authentication with server-side sessions
- Adding file uploads requiring storage
- Adding WebSocket features (increases resource requirements)

## Troubleshooting

### Port Already in Use

If you see an error about the port being in use:

```bash
# Use a different port
PORT=8081 go run .
```

### CGO Errors

If you encounter CGO-related errors:

```bash
# Ensure CGO is enabled
CGO_ENABLED=1 go build .
```

On Linux/macOS, install required dependencies:

```bash
# Debian/Ubuntu
sudo apt-get install gcc musl-dev sqlite3

# macOS
brew install gcc sqlite
```

### Template Parse Errors

If templates fail to parse (server.go:61-66):
1. Check HTML syntax in `server/pages/` files
2. Ensure all `{{define}}` blocks have matching `{{end}}`
3. Verify the layout structure is correct

### Docker Build Fails

Ensure Docker daemon is running:

```bash
docker ps
```

Try rebuilding without cache:

```bash
docker build --no-cache -f config/Dockerfile -t todo-tasker .
```

### External Docker Resources Required

The `compose.yml` requires external Docker resources. Create them before running:

```bash
# Create external network
docker network create deploy_net

# Create external volume
docker volume create deploy_sqlite_dir
```

## Deployment

### Automated Deployment with Ansible

The project includes Ansible playbooks for automated VPS deployment:

1. GitHub Actions workflow triggers on push to `main` branch
2. Ansible playbook runs on your VPS (`ansible/playbook.yml`)
3. Code is pulled from repository
4. Containers are stopped, rebuilt, and restarted

**Required GitHub Secrets:**
- `VPS_IP` - Your server's IP address
- `ANSIBLE_SSH_USER` - SSH username
- `VPS_SSH_PRIVATE_KEY` - SSH private key

### Manual Deployment

On your VPS:

```bash
# Pull latest code
git pull origin main

# Stop existing containers
make stop

# Rebuild and start
make run
```

## Next Steps

- Customize templates in `server/pages/` to match your branding
- Add new components in `server/pages/components/`
- Implement additional routes in `server/server.go` and `server/routes.go`
- Write tests for new features in `tests/`
- Configure your VPS for production deployment
- Set up SSL/TLS certificates for HTTPS
- Implement user authentication (client-side recommended)
- Add data export/import functionality

## Support

For issues or questions:
- Open an issue on GitHub
- Review test files for implementation examples
- Check the server code for API details

## License

See the LICENSE file for details.
