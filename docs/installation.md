# Installation Guide

Complete installation instructions for Todo-Tasker, from cloning the repository to verifying your setup.

## Quick Installation

```bash
# Clone the repository
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker

# Download dependencies
go mod download

# Verify installation
go mod verify

# Run the application
make dev
```

That's it! Access at `http://localhost:8080`

## Detailed Installation Steps

### Step 1: Clone the Repository

Using HTTPS (recommended):
```bash
git clone https://github.com/duet-ink/Todo-Tasker.git
```

Using SSH (if you have SSH keys configured):
```bash
git clone git@github.com:HarudaySharma/Todo-Tasker.git
```

### Step 2: Navigate to Project Directory

```bash
cd Todo-Tasker
```

### Step 3: Verify Repository Structure

Check that you have the essential files:
```bash
ls -la
```

You should see:
- `app.go` - Main application entry point
- `go.mod` - Go module definition
- `Makefile` - Build commands
- `config/` - Configuration directory
- `server/` - HTTP server code
- `docs/` - Documentation
- `tests/` - Test files

### Step 4: Install Go Dependencies

Download all required Go modules:
```bash
go mod download
```

This will download any dependencies specified in `go.mod`.

### Step 5: Verify Dependencies

Ensure all dependencies are correctly downloaded:
```bash
go mod verify
```

Expected output:
```
all modules verified
```

### Step 6: Clean Up Unused Dependencies (Optional)

Remove any unused dependencies:
```bash
go mod tidy
```

### Step 7: Test Build

Verify the application compiles successfully:
```bash
go build -o todo-tasker .
```

This creates a binary named `todo-tasker`. You can delete it after verification:
```bash
rm todo-tasker
```

## Docker Installation (Optional)

If you plan to use Docker for deployment:

### Create Docker Resources

```bash
# Create external network
docker network create deploy_net

# Create external volume (for future features)
docker volume create deploy_sqlite_dir
```

### Verify Docker Setup

```bash
# Check Docker is running
docker ps

# Test build
docker build -f config/Dockerfile -t todo-tasker:test .

# Clean up test image
docker rmi todo-tasker:test
```

## Environment Configuration

### Option 1: Environment Variables

Set environment variables directly:
```bash
export PORT=8080
export ADMIN_PORT=4660
```

### Option 2: Create app.env File

Create an environment file in the project root:
```bash
cat > app.env << 'EOF'
PORT=8080
ADMIN_PORT=4660
EOF
```

**Note**: `app.env` is in `.gitignore` and won't be committed to version control.

## Verification Checklist

Before running the application, verify:

- [ ] Repository cloned successfully
- [ ] In project directory (`Todo-Tasker/`)
- [ ] Go dependencies downloaded (`go mod download`)
- [ ] Dependencies verified (`go mod verify`)
- [ ] Application builds without errors (`go build .`)
- [ ] Docker resources created (if using Docker)
- [ ] Environment configured (optional)

## Platform-Specific Notes

### Linux

**Ubuntu/Debian:**
```bash
# Install build tools
sudo apt-get update
sudo apt-get install make

# Clone and install
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker
go mod download
```

**Fedora/RHEL:**
```bash
# Install build tools
sudo dnf install make

# Clone and install
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker
go mod download
```

### macOS

```bash
# Install Xcode Command Line Tools (includes make)
xcode-select --install

# Install Go via Homebrew (if not already installed)
brew install go

# Clone and install
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker
go mod download
```

### Windows

**Using WSL2 (Recommended):**
```bash
# Install WSL2 with Ubuntu
wsl --install

# Follow Linux installation steps in WSL
```

**Native Windows:**
```powershell
# Clone repository
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker

# Install dependencies
go mod download
```

## Troubleshooting Installation

### "go.mod: no such file or directory"

Ensure you're in the correct directory:
```bash
pwd
# Should show: /path/to/Todo-Tasker

cd Todo-Tasker
ls go.mod
```

### "permission denied" Errors

Fix ownership of the directory:
```bash
sudo chown -R $USER:$USER .
```

Or clone to your home directory:
```bash
cd ~
git clone https://github.com/duet-ink/Todo-Tasker.git
```

### Dependency Download Fails

**Check internet connection**, then try:

```bash
# Use direct mode (bypass proxy)
GOPROXY=direct go mod download

# Or use a different proxy
GOPROXY=https://goproxy.io,direct go mod download

# Check proxy settings
go env GOPROXY
```

### Build Errors

**Error**: Build fails with module errors

**Solution**: Clean and rebuild:
```bash
go clean -modcache
go mod download
go mod verify
go build .
```

### Build Fails with "cannot find package"

Ensure all dependencies are downloaded:
```bash
go mod download
go mod verify
go mod tidy
```

### Git Clone Fails

**Check network connection**, then try:

```bash
# Try HTTPS instead of SSH
git clone https://github.com/duet-ink/Todo-Tasker.git

# Or SSH instead of HTTPS
git clone git@github.com:HarudaySharma/Todo-Tasker.git

# Check Git installation
git --version
```

### Docker Errors

**Error**: `Cannot connect to the Docker daemon`

**Solution**: Start Docker:
```bash
# Linux
sudo systemctl start docker

# macOS/Windows
# Start Docker Desktop application
```

**Error**: `Error response from daemon: network not found`

**Solution**: Create required networks:
```bash
docker network create deploy_net
docker volume create deploy_sqlite_dir
```

## Updating Dependencies

To update to the latest compatible versions:

```bash
# Update all dependencies
go get -u ./...

# Tidy up
go mod tidy

# Verify
go mod verify

# Test everything still works
make test
```

## Uninstallation

To completely remove Todo-Tasker:

```bash
# Stop running containers (if any)
make stop

# Remove Docker resources
docker network rm deploy_net
docker volume rm deploy_sqlite_dir

# Remove project directory
cd ..
rm -rf Todo-Tasker
```

## Next Steps

After successful installation:

1. **[Run the application](get-started.md#running-the-application)** - Start the dev server
2. **[Read the developer guide](developer.md)** - Learn clean code practices
3. **[Check the todo list](todo.md)** - Find tasks to work on
4. **[Review architecture](architecture.md)** - Understand the system design

## Installation Support

If you encounter issues not covered here:

1. Check the [troubleshooting section](get-started.md#troubleshooting)
2. Search [existing issues](https://github.com/duet-ink/Todo-Tasker/issues)
3. Create a [new issue](https://github.com/duet-ink/Todo-Tasker/issues/new) with:
   - Your OS and version
   - Go version (`go version`)
   - Error messages
   - Steps to reproduce

---

**[← Back to Get Started](get-started.md)**
