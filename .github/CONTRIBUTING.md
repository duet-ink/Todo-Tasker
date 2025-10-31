# Contributing to Todo-Tasker

Thank you for your interest in contributing to Todo-Tasker! This document provides guidelines and instructions for contributing to the project.

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [Getting Started](#getting-started)
3. [How to Contribute](#how-to-contribute)
4. [Development Workflow](#development-workflow)
5. [Coding Standards](#coding-standards)
6. [Testing Guidelines](#testing-guidelines)
7. [Pull Request Process](#pull-request-process)
8. [Issue Guidelines](#issue-guidelines)
9. [Community](#community)

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [support@duet.ink](mailto:support@duet.ink).

## Getting Started

### Prerequisites

Before you start contributing, ensure you have:

1. Go 1.23.3 or higher installed
2. Git for version control
3. A GitHub account
4. (Optional) Make for using Makefile commands
5. (Optional) Docker for containerized development

See [docs/get-started.md](docs/get-started.md) for detailed setup instructions.

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR-USERNAME/Todo-Tasker.git
   cd Todo-Tasker
   ```
3. Add upstream remote:
   ```bash
   git remote add upstream https://github.com/duet-ink/Todo-Tasker.git
   ```
4. Install dependencies:
   ```bash
   go mod download
   ```

### Set Up Development Environment

```bash
# Run the application
make dev

# Run tests
make test

# Access at http://localhost:8080
```

## How to Contribute

### Ways to Contribute

- **Report bugs**: Create an issue with reproduction steps
- **Suggest features**: Open an issue to discuss new ideas
- **Fix bugs**: Pick up issues labeled `bug` or `good first issue`
- **Implement features**: Check [docs/todo.md](docs/todo.md) for planned features
- **Improve documentation**: Fix typos, add examples, clarify instructions
- **Write tests**: Increase test coverage
- **Review pull requests**: Provide constructive feedback

### Finding Tasks

1. Check [docs/todo.md](docs/todo.md) for planned features and tasks
2. Browse [issues labeled "good first issue"](https://github.com/duet-ink/Todo-Tasker/labels/good%20first%20issue)
3. Look for [issues labeled "help wanted"](https://github.com/duet-ink/Todo-Tasker/labels/help%20wanted)
4. Ask in issues or discussions what needs attention

## Development Workflow

### 1. Create a Feature Branch

```bash
# Update your local main branch
git checkout main
git pull upstream main

# Create feature branch
git checkout -b feature/your-feature-name
```

Branch naming conventions:
- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation updates
- `refactor/` - Code refactoring
- `test/` - Test additions/updates

### 2. Make Your Changes

- Follow the [coding standards](#coding-standards)
- Write clear, self-documenting code
- Add tests for new functionality
- Update documentation as needed

### 3. Commit Your Changes

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```bash
git add .
git commit -m "feat: add WebSocket sync endpoint"
```

Commit message format:
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style (formatting, missing semicolons, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(server): add WebSocket sync endpoint

Implements P2P device sync via WebSocket.
Includes session management and message routing.

Closes #42
```

```
fix(middleware): correct CORS header for components

The Access-Control-Allow-Origin header was not set
correctly for component POST requests.
```

### 4. Keep Your Branch Updated

```bash
# Fetch latest changes from upstream
git fetch upstream

# Rebase your branch
git rebase upstream/main

# Resolve conflicts if any
```

### 5. Run Tests

```bash
# Run all tests
make test

# Run with race detector
go test -race ./tests/

# Check coverage
go test -cover ./tests/

# Lint code (if golangci-lint installed)
golangci-lint run
```

### 6. Push Your Branch

```bash
git push origin feature/your-feature-name
```

## Coding Standards

### Go Code Standards

Follow the guidelines in [docs/developer.md](docs/developer.md).

**Key principles:**
- Write clean, readable code
- Keep functions small and focused
- Handle errors explicitly
- Use meaningful variable names
- Document public functions
- Follow Go conventions

**Format and vet your code:**
```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run linter
golangci-lint run
```

### Frontend Standards

**HTML:**
- Use semantic HTML5 elements
- Include ARIA labels for accessibility
- Follow HTMX best practices

**JavaScript/Alpine.js:**
- Keep components small and focused
- Use clear, descriptive names
- Handle errors gracefully
- Avoid global state pollution

**CSS/Tailwind:**
- Use utility classes for one-offs
- Create component classes for repeated patterns
- Follow mobile-first approach

## Testing Guidelines

### Writing Tests

**Go tests:**
```go
// tests/feature_test.go
func TestFeatureName(t *testing.T) {
    // Arrange
    input := setupTestInput()

    // Act
    result := performOperation(input)

    // Assert
    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}
```

**Test organization:**
- Place tests in `tests/` directory
- Use table-driven tests for multiple cases
- Test both happy path and error conditions
- Mock external dependencies

### Running Tests

```bash
# All tests
go test ./tests/ -v

# Specific test file
go test ./tests/server_test.go -v

# With coverage
go test -cover ./tests/

# Generate coverage report
go test -coverprofile=coverage.out ./tests/
go tool cover -html=coverage.out
```

### Test Coverage

- Aim for meaningful coverage, not just high percentage
- Focus on critical business logic
- Test error conditions and edge cases
- Don't test external libraries

## Pull Request Process

### Before Submitting

- [ ] Code follows project conventions
- [ ] All tests pass locally
- [ ] New tests added for new functionality
- [ ] Documentation updated if needed
- [ ] Commit messages follow conventions
- [ ] No merge conflicts with main branch
- [ ] Code is formatted (`go fmt`)
- [ ] Code passes vet (`go vet`)

### Submitting a Pull Request

1. **Push your branch** to your fork
2. **Create PR** on GitHub
3. **Fill out PR template** with:
   - Description of changes
   - Related issue numbers
   - Testing performed
   - Screenshots (if UI changes)

### PR Title Format

```
<type>: <description>
```

Examples:
```
feat: add WebSocket sync endpoint
fix: correct CORS headers in middleware
docs: update installation guide
```

### PR Description Template

```markdown
## Description
Brief description of what this PR does.

## Related Issues
Closes #123
Related to #456

## Changes Made
- Change 1
- Change 2
- Change 3

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing performed

## Screenshots (if applicable)
Add screenshots for UI changes

## Checklist
- [ ] Code follows project conventions
- [ ] Tests pass locally
- [ ] Documentation updated
- [ ] Commit messages follow conventions
```

### Code Review

- Respond to feedback promptly and politely
- Make requested changes in new commits
- Don't force-push after review starts (unless requested)
- Mark conversations as resolved when addressed
- Request re-review after making changes

### Merging

- Maintainers will merge approved PRs
- PRs are typically squash-merged to keep history clean
- Delete your branch after merge

## Issue Guidelines

### Before Creating an Issue

1. Search existing issues to avoid duplicates
2. Check [docs/todo.md](docs/todo.md) - feature might be planned
3. Try to reproduce the bug with latest version

### Bug Reports

Use the bug report template:

```markdown
**Describe the bug**
A clear description of what the bug is.

**To Reproduce**
Steps to reproduce:
1. Go to '...'
2. Click on '...'
3. See error

**Expected behavior**
What you expected to happen.

**Actual behavior**
What actually happened.

**Environment:**
- OS: [e.g., Ubuntu 22.04]
- Go version: [e.g., 1.23.3]
- Browser: [e.g., Chrome 120]

**Screenshots**
If applicable, add screenshots.

**Additional context**
Any other relevant information.
```

### Feature Requests

Use the feature request template:

```markdown
**Is your feature request related to a problem?**
A clear description of the problem.

**Describe the solution you'd like**
A clear description of what you want to happen.

**Describe alternatives you've considered**
Other solutions or features you've considered.

**Additional context**
Any other context or screenshots.
```

### Issue Labels

- `bug` - Something isn't working
- `feature` - New feature or enhancement
- `documentation` - Documentation improvements
- `good first issue` - Good for newcomers
- `help wanted` - Extra attention needed
- `question` - Further information requested
- `wontfix` - This will not be worked on

## Community

### Communication Channels

- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: Questions, ideas, general discussion
- **Pull Requests**: Code review and collaboration
- **Email**: support@duet.ink

### Getting Help

- Check [docs/get-started.md](docs/get-started.md) for setup help
- Review [docs/developer.md](docs/developer.md) for coding guidelines
- Search existing issues for similar problems
- Ask questions in GitHub Discussions
- Tag maintainers in issues if stuck

### Recognition

Contributors are recognized in several ways:
- Listed in GitHub contributors
- Mentioned in release notes
- Added to CHANGELOG for significant contributions
- Acknowledged in commit messages

## Project Structure

For a detailed overview of the project structure, file organization, and architecture, see:

**[docs/get-started.md - Project Structure](../docs/get-started.md#project-structure)**

Key directories:
- `.github/` - Community files and CI/CD workflows
- `config/` - Configuration and Docker setup
- `server/` - HTTP server, routes, templates, and assets
- `tests/` - Test files
- `docs/` - Documentation
- `ansible/` - Deployment automation

## Development Tips

### Running with Hot Reload

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with auto-reload
air
```

### Debugging

```bash
# Run with verbose logging
go run . -v

# Use delve debugger
dlv debug
```

### Docker Development

```bash
# Build and run with Docker
make run

# View logs
docker compose -f config/compose.yml logs -f

# Stop containers
make stop
```

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (see LICENSE file).

## Questions?

If you have questions about contributing:
1. Check this guide and [docs/developer.md](docs/developer.md)
2. Search existing issues and discussions
3. Open a discussion or issue with the "question" label
4. Contact maintainers at [support@duet.ink](mailto:support@duet.ink)

---

Thank you for contributing to Todo-Tasker! Your efforts help make this project better for everyone.
