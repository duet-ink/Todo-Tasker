# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- SQLite WASM database initialization
- Todo CRUD operations (Create, Read, Update, Delete)
- Device-to-device sync via WebSocket
- QR code generation and scanning for sync
- PWA functionality with service worker
- Offline support
- Data export/import functionality

## [0.0.7] - 2025-10-31

### Added
- Community files now properly organized in `.github/` directory
  - CODE_OF_CONDUCT.md moved to `.github/`
  - CONTRIBUTING.md moved to `.github/`
  - CHANGELOG.md moved to `.github/`
- Created comprehensive docs/architecture.md (system design)
- Support email contact information throughout documentation

### Changed
- **BREAKING**: Updated repository URL from `github.com/HarudaySharma/Todo-Tasker` to `github.com/duet-ink/Todo-Tasker`
- Updated support email from `harudaysharma@gmail.com` to `support@duet.ink`
- Streamlined documentation structure
  - Consolidated getting started guide (removed redundant separate files)
  - Removed docs/docker-deployment.md (content merged into get-started.md)
  - Removed docs/prerequisites.md (content merged into get-started.md)
  - Removed docs/running-locally.md (content merged into get-started.md)
  - Removed docs/usage.md (content merged into get-started.md)
  - Removed docs/project-structure.md (content merged into get-started.md)
- Renamed `docs/todos.md` to `docs/todo.md` for consistency
- Enhanced README.md with clearer project links section
- Updated all internal documentation links to reflect new structure
- Improved docs/get-started.md with all-in-one comprehensive guide
- Enhanced docs/installation.md with platform-specific instructions

### Fixed
- Corrected all broken documentation cross-references
- Fixed sync-architecture.md references (now points to architecture.md)
- Updated GitHub issue and repository links throughout all documentation

## [0.0.6] - 2025-10-31

### Added
- Comprehensive documentation structure
- Developer guide with clean code practices (docs/developer.md)
- Contributing guidelines (CONTRIBUTING.md)
- Code of Conduct (CODE_OF_CONDUCT.md)
- Consolidated getting started guide (docs/get-started.md)
- Detailed installation guide (docs/installation.md)
- Feature roadmap and task list (docs/todo.md)

### Changed
- Restructured documentation for better organization
- Updated README.md with project overview and quick links
- Simplified documentation structure by consolidating guides

## [0.0.5] - 2025-10-30

### Added
- Sync architecture documentation (docs/sync-architecture.md)
- Detailed sync protocol specification
- WebSocket-based P2P sync design

### Changed
- Updated project documentation
- Enhanced sync planning documentation

## [0.0.4] - 2025-10-29

### Added
- Initial documentation structure
- Getting started guide
- Project structure documentation

### Fixed
- Makefile reference to deploy.yml (corrected to compose.yml)

## [0.0.3] - 2025-10-28

### Added
- Component endpoint at POST /c/{name} for dynamic component loading
- HTMX component loading system
- Additional components: assistent, dashboard, settings, sidebar

### Changed
- Enhanced component rendering system
- Improved error handling for component requests

## [0.0.2] - 2025-10-27

### Added
- Basic HTML templates (layout, index, admin, error)
- HTMX integration for dynamic content loading
- Alpine.js integration for client-side reactivity
- DaisyUI and Tailwind CSS for styling
- SQLite WASM files for client-side database
- Favicon and logo assets
- PWA manifest.json

### Changed
- Improved template structure with layout system
- Enhanced frontend component organization

## [0.0.1] - 2025-10-26

### Added
- Initial project setup
- Go module initialization (go.mod, go.sum)
- Main application entry point (app.go)
- Two-server architecture (main server and admin server)
- Configuration package (config/config.go)
- Environment variable support (PORT, ADMIN_PORT)
- JSON structured logging with slog
- HTTP server implementation (server/server.go)
- Basic route handlers (server/routes.go)
- Template embedding with //go:embed
- Static asset serving
- Middleware structure (database, redis, security)
- Test infrastructure (tests/)
- Docker configuration (config/Dockerfile, config/compose.yml)
- Makefile with dev, run, stop, test commands
- Ansible deployment automation (ansible/playbook.yml)
- GitHub Actions CI/CD workflow (.github/workflows/ansible.yml)
- Basic project documentation (README.md, CLAUDE.md)

### Infrastructure
- Docker Compose setup with external networks and volumes
- Automated VPS deployment via Ansible
- GitHub Actions integration for CI/CD

---

## Version History

- **0.0.7** - Documentation restructuring and organization consolidation
- **0.0.6** - Documentation overhaul and developer guidelines
- **0.0.5** - Sync architecture documentation
- **0.0.4** - Initial documentation
- **0.0.3** - Component system implementation
- **0.0.2** - Frontend integration
- **0.0.1** - Initial project setup

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for information on how to contribute to this project.

## Links

- [GitHub Repository](https://github.com/duet-ink/Todo-Tasker)
- [Issue Tracker](https://github.com/duet-ink/Todo-Tasker/issues)
- [Documentation](docs/get-started.md)
- [Support Email](mailto:support@duet.ink)
