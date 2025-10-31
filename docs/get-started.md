# Getting Started with Todo-Tasker

A comprehensive guide to set up, develop, and deploy your cost-efficient todo management application.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Running the Application](#running-the-application)
5. [Project Structure](#project-structure)
6. [Development Workflow](#development-workflow)
7. [Testing](#testing)
8. [Deployment](#deployment)
9. [Troubleshooting](#troubleshooting)

## Project Overview

Todo-Tasker is a lightweight, scalable todo application built with a unique architecture that eliminates traditional database costs. The entire application can serve thousands of users on minimal infrastructure with hosting costs under $10/month.

### Why Todo-Tasker?

- **Zero Database Costs**: SQLite WASM runs entirely in the browser (IndexedDB)
- **Stateless Backend**: Go server only serves HTML and assets, enabling cheap horizontal scaling
- **Single Binary Deployment**: All assets embedded at compile time for simplified deployment
- **Privacy-Focused**: All user data stays on their devices, never stored on servers
- **Offline-First PWA**: Full functionality without internet connection
- **Device Sync**: P2P sync via WebSocket without server-side data storage

### Tech Stack

**Backend:**
- Go 1.23.3+ with standard library HTTP server
- `//go:embed` for asset bundling
- Pure Go implementation (no CGO required)
- Native Go tools and standard library preferred

**Frontend:**
- HTMX 2.0.7 for dynamic HTML loading
- Alpine.js 3.x for client-side reactivity
- DaisyUI 5 + Tailwind CSS 4 for styling
- SQLite WASM for client-side database
- Tabler Icons for UI elements

**Infrastructure:**
- Docker & Docker Compose
- Ansible for deployment automation
- GitHub Actions for CI/CD

## Prerequisites

Before you begin, ensure you have the following installed:

### Required Software

**Go 1.23.3 or Higher**
```bash
# Download from https://golang.org/dl/
# Verify installation
go version
```

**Git**
```bash
# Download from https://git-scm.com/downloads
# Verify installation
git --version
```

**Note:** This project uses pure Go and SQLite WASM (client-side in browser), so no C compiler or CGO is required.

### Optional Tools

**Make** (Recommended)
```bash
# Linux
sudo apt-get install build-essential

# macOS (usually pre-installed)
xcode-select --install
```

**Docker** (For containerized deployment)
```bash
# Download from https://www.docker.com/get-started
docker --version
docker compose version
```

**Air** (For hot reload during development)
```bash
go install github.com/cosmtrek/air@latest
```

## Installation

See [installation.md](installation.md) for detailed installation instructions.

### Quick Install

```bash
# Clone the repository
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker

# Download dependencies
go mod download

# Verify installation
go mod verify
```

## Running the Application

### Development Mode (Recommended)

Using Make:
```bash
make dev
```

Or manually:
```bash
PORT=8080 ADMIN_PORT=4660 go run .
```

Access at: `http://localhost:8080`

### Production Mode

With Docker:
```bash
make run
```

Or build and run binary:
```bash
go build -ldflags="-s -w" -o todo-tasker .
./todo-tasker
```

### With Hot Reload

```bash
air
```

## Project Structure

```
Todo-Tasker/
├── app.go                    # Main application entry point
├── go.mod, go.sum           # Go module files
├── Makefile                 # Build commands
├── README.md                # Project overview
├── CLAUDE.md                # Claude Code instructions
│
├── .github/                 # GitHub configuration
│   ├── CONTRIBUTING.md     # Contribution guidelines
│   ├── CODE_OF_CONDUCT.md  # Code of conduct
│   ├── CHANGELOG.md        # Version history
│   └── workflows/          # CI/CD pipelines
│
├── config/                  # Configuration
│   ├── config.go           # Environment config
│   ├── Dockerfile          # Docker configuration
│   └── compose.yml         # Docker Compose setup
│
├── server/                  # HTTP server
│   ├── server.go           # Server setup
│   ├── routes.go           # HTTP handlers
│   ├── pages/              # HTML templates (embedded)
│   │   ├── layout.html
│   │   ├── index.html
│   │   └── components/     # HTMX components
│   ├── assets/             # Static files (embedded)
│   │   ├── styles.css
│   │   ├── database/       # SQLite WASM files
│   │   └── images/
│   └── middleware/         # HTTP middleware
│
├── tests/                   # Test files
├── docs/                    # Documentation
│   ├── get-started.md      # This file
│   ├── installation.md     # Installation guide
│   ├── developer.md        # Clean code practices
│   ├── architecture.md     # System design
│   └── todo.md             # Feature roadmap
│
└── ansible/                 # Deployment automation
```

### Key Files

- **app.go:15** - Main entry point, launches two concurrent servers
- **server/server.go:80** - Route registration and template management
- **server/routes.go:8** - HTTP handler implementations
- **config/config.go** - Environment configuration and logging

## Development Workflow

### 1. Create a New Branch

```bash
git checkout -b feature/your-feature-name
```

### 2. Make Changes

Follow clean code practices outlined in [developer.md](developer.md).

### 3. Test Your Changes

```bash
# Run tests
make test

# Run with race detector
go test -race ./tests/

# Check coverage
go test -cover ./tests/
```

### 4. Format and Lint

```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Lint (if golangci-lint installed)
golangci-lint run
```

### 5. Commit Changes

Follow conventional commits:
```bash
git add .
git commit -m "feat: add new feature description"
```

### 6. Push and Create PR

```bash
git push origin feature/your-feature-name
```

Create a pull request on GitHub.

## Testing

### Run All Tests

```bash
make test
# or
go test ./tests/ -v
```

### Test Categories

```bash
# Configuration tests
go test ./tests/config_test.go -v

# Server tests
go test ./tests/server_test.go -v

# Middleware tests
go test ./tests/middleware_test.go -v

# Component tests
go test ./tests/components_test.go -v
```

### Coverage Report

```bash
go test -cover ./tests/
go test -coverprofile=coverage.out ./tests/
go tool cover -html=coverage.out
```

## Deployment

### Docker Deployment

```bash
# Build and run
make run

# Stop containers
make stop

# View logs
docker compose -f config/compose.yml logs -f
```

### Manual VPS Deployment

```bash
# On your VPS
git pull origin main
make stop
make run
```

### Automated Deployment

Push to `main` branch triggers GitHub Actions workflow:
1. Ansible playbook runs on VPS
2. Pulls latest code
3. Rebuilds containers
4. Restarts services

**Required GitHub Secrets:**
- `VPS_IP`
- `ANSIBLE_SSH_USER`
- `VPS_SSH_PRIVATE_KEY`

## Troubleshooting

### Port Already in Use

```bash
# Use different port
PORT=8081 go run .
```

### Build Errors

If you encounter build errors:

```bash
# Ensure Go version is correct
go version

# Clean build cache
go clean -cache

# Rebuild
go build .
```

### Docker Build Fails

```bash
# Clear cache
docker builder prune -a

# Rebuild
docker build --no-cache -f config/Dockerfile -t todo-tasker .
```

### Templates Not Updating

Remember: templates are embedded at compile time.
```bash
# Stop server (Ctrl+C)
# Rebuild and restart
go run .
```

### External Docker Resources

```bash
# Create required resources
docker network create deploy_net
docker volume create deploy_sqlite_dir
```

## Available Routes

### Public Routes

- `/` - Main application
- `/404/` - 404 error page
- `/error/` - Error page
- `POST /c/{name}` - Component endpoint
- `/assets/*` - Static assets

### Admin Routes

- `/admin/` - Admin interface

### Components

Access via `POST /c/{name}`:
- navbar, sidebar, footer
- dashboard, todo, settings
- login, assistent

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Main server port | `80` |
| `ADMIN_PORT` | Admin server port | `4657` |

### Create app.env

```bash
cat > app.env << 'EOF'
PORT=8080
ADMIN_PORT=4660
EOF
```

## Next Steps

1. **[Read Developer Guide](developer.md)** - Learn clean code practices
2. **[Review Contributing Guidelines](../.github/CONTRIBUTING.md)** - Contribution workflow
3. **[Check Todo List](todo.md)** - Available tasks
4. **[Understand Architecture](architecture.md)** - System design
5. **[Read Code of Conduct](../.github/CODE_OF_CONDUCT.md)** - Community guidelines

## Resources

- **[GitHub Repository](https://github.com/duet-ink/Todo-Tasker)**
- **[Issue Tracker](https://github.com/duet-ink/Todo-Tasker/issues)**
- **[Changelog](../.github/CHANGELOG.md)**

## Support

- **Found a bug?** [Create an issue](https://github.com/duet-ink/Todo-Tasker/issues)
- **Feature request?** Check [todo.md](todo.md) first, then open an issue
- **Need help?** Open an issue with the "question" label
- **Email support**: support@duet.ink

---

**Current Version:** v0.0.7
**Last Updated:** 2025-10-31
