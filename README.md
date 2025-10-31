# Todo-Tasker

A cost-efficient, scalable multi-user todo management application designed to run entirely on client-side storage with minimal server overhead.

## Project Goal

**Build a production-ready todo application with total hosting costs under $10/month**, serving thousands of users on minimal infrastructure.

## Key Features

- **Zero Database Costs**: SQLite WASM runs entirely in the browser (IndexedDB)
- **Stateless Backend**: Lightweight Go server for serving HTML and assets only
- **Single Binary Deployment**: All assets embedded at compile time
- **Offline-First PWA**: Full functionality without internet connection
- **Device-to-Device Sync**: WebSocket-based P2P sync without storing user data
- **Privacy-Focused**: All user data stays on their devices

## Tech Stack

**Backend:**
- Go 1.23.3+
- Standard library HTTP server
- Embedded assets (`//go:embed`)

**Frontend:**
- HTMX 2.0.7 - Dynamic HTML loading
- Alpine.js 3.x - Client-side reactivity
- DaisyUI 5 - UI components
- Tailwind CSS 4 - Utility-first styling
- SQLite WASM - Client-side database
- Tabler Icons - Icon library

**Infrastructure:**
- Docker & Docker Compose
- Ansible for deployment automation
- GitHub Actions CI/CD

## Getting Started

For detailed setup instructions, architecture documentation, and development guides, see:

**[📚 Get Started Guide →](docs/get-started.md)**

Quick start:

```bash
git clone https://github.com/duet-ink/Todo-Tasker.git
cd Todo-Tasker
make dev
```

Access at `http://localhost:8080`

## Documentation

- **[Get Started](docs/get-started.md)** - Complete setup and development guide
- **[Architecture](docs/architecture.md)** - System design and architecture
- **[Todo List](docs/todo.md)** - Feature roadmap and development tasks

## Contributing

We welcome contributions! Please:

1. Check [docs/todo.md](docs/todo.md) for available tasks and feature requests
2. Create an issue if you encounter bugs or need support
3. Follow the architecture guidelines in [Code of Conduct](.github/CODE_OF_CONDUCT.md)
4. Read [Contributing Guidelines](.github/CONTRIBUTING.md)
5. Run tests before submitting PRs: `make test`

## Cost Breakdown

**Target: <$10/month for thousands of users**

- VPS hosting: $5-6/month (1GB RAM)
- Domain: ~$1/month
- Database: $0 (client-side only)
- CDN: $0 (free public CDNs)
- External APIs: $0 (none required)

## Support

- **Issues & Bugs**: [Create an issue](https://github.com/duet-ink/Todo-Tasker/issues)
- **Feature Requests**: Check [docs/todo.md](docs/todo.md) or open an issue
- **Support Email**: support@duet.ink
- **Questions**: Review documentation or open a discussion

## Project Links

- **[Contributing Guidelines](.github/CONTRIBUTING.md)** - How to contribute
- **[Code of Conduct](.github/CODE_OF_CONDUCT.md)** - Community guidelines
- **[Changelog](.github/CHANGELOG.md)** - Version history

## License

See [LICENSE](LICENSE) file for details.

---

**Current Version:** v0.0.7
