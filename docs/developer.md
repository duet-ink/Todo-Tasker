# Developer Guide

Essential practices for writing clean, maintainable code in Todo-Tasker.

## Table of Contents

1. [Project Philosophy](#project-philosophy)
2. [Clean Code Principles](#clean-code-principles)
3. [Go Best Practices](#go-best-practices)
4. [Frontend Best Practices](#frontend-best-practices)
5. [Code Organization](#code-organization)
6. [Naming Conventions](#naming-conventions)
7. [Error Handling](#error-handling)
8. [Testing Guidelines](#testing-guidelines)
9. [Documentation Standards](#documentation-standards)
10. [Git Workflow](#git-workflow)
11. [Code Review Checklist](#code-review-checklist)

## Project Philosophy

### Native Tools & Pure Go

Todo-Tasker is built with **pure Go** and emphasizes using native tools wherever possible.

**Core Principles:**

1. **No CGO Dependencies**
   - Pure Go implementation (CGO_ENABLED=0)
   - No C compiler required
   - Easier builds and cross-compilation
   - Better portability across platforms

2. **Prefer Standard Library**
   - Use Go's standard library over external packages when practical
   - Reduces dependency bloat
   - Better long-term maintainability
   - Faster compile times

3. **Client-Side Processing**
   - SQLite WASM runs entirely in the browser
   - No server-side database required
   - Reduces server complexity and costs
   - Enhances user privacy

4. **Minimal External Dependencies**
   - Only add dependencies when truly necessary
   - Evaluate alternatives in standard library first
   - Consider maintenance burden of each dependency
   - Keep `go.mod` lean

**Benefits:**
- ✅ Faster builds (no C compilation)
- ✅ Static binary generation (single executable)
- ✅ Cross-platform compilation (easy GOOS/GOARCH)
- ✅ Smaller attack surface
- ✅ Simpler deployment
- ✅ Better long-term maintainability

**When to Add Dependencies:**
- Standard library doesn't provide the functionality
- Significant complexity reduction
- Well-maintained, popular package
- Security-critical functionality (crypto libraries)

## Clean Code Principles

### 1. Write Code for Humans

Code is read far more often than it's written. Prioritize readability over cleverness.

**Bad:**
```go
func p(d int) int {
    if d < 0 { return -d }
    return d
}
```

**Good:**
```go
func calculatePriority(daysOverdue int) int {
    if daysOverdue < 0 {
        return 0
    }
    return daysOverdue
}
```

### 2. Single Responsibility Principle

Each function, struct, or module should have one clear purpose.

**Bad:**
```go
func processUserAndSendEmail(user User) error {
    // Validates user
    // Updates database
    // Sends email
    // Logs activity
    // All in one function!
}
```

**Good:**
```go
func validateUser(user User) error { /* ... */ }
func updateUserInDatabase(user User) error { /* ... */ }
func sendWelcomeEmail(user User) error { /* ... */ }
func logUserActivity(user User, action string) { /* ... */ }
```

### 3. DRY (Don't Repeat Yourself)

Eliminate code duplication through abstraction.

**Bad:**
```go
// Repeated validation logic
if todo.Title == "" {
    return errors.New("title is required")
}
if len(todo.Title) > 200 {
    return errors.New("title too long")
}
// ... repeated in multiple places
```

**Good:**
```go
func validateTodoTitle(title string) error {
    if title == "" {
        return errors.New("title is required")
    }
    if len(title) > 200 {
        return errors.New("title too long")
    }
    return nil
}
```

### 4. KISS (Keep It Simple, Stupid)

Simplicity should be a key goal. Avoid unnecessary complexity.

**Bad:**
```go
func getTodoStatus(completed bool, deleted bool, archived bool) string {
    if completed && !deleted && !archived {
        return "done"
    } else if !completed && !deleted && !archived {
        return "active"
    } else if deleted {
        return "deleted"
    } else if archived {
        return "archived"
    }
    return "unknown"
}
```

**Good:**
```go
func getTodoStatus(todo Todo) string {
    switch {
    case todo.Deleted:
        return "deleted"
    case todo.Archived:
        return "archived"
    case todo.Completed:
        return "done"
    default:
        return "active"
    }
}
```

### 5. YAGNI (You Aren't Gonna Need It)

Don't add functionality until it's necessary.

**Bad:**
```go
// Adding features "just in case"
type Todo struct {
    ID          string
    Title       string
    // Future features we might need?
    Color       string
    Icon        string
    ParentID    string
    ChildIDs    []string
    Collaborators []User
    // ...
}
```

**Good:**
```go
// Only what we need now
type Todo struct {
    ID          string
    Title       string
    Description string
    Completed   bool
    Priority    int
    CreatedAt   time.Time
}
```

## Go Best Practices

### Package Organization

```go
// Good package structure
package server

import (
    "net/http"
    "github.com/HarudaySharma/Todo-Tasker/config"
)

// Exported types and functions (capital letter)
type Server struct {
    // ...
}

func New() *Server {
    // ...
}

// Unexported helpers (lowercase letter)
func parseTemplate(name string) error {
    // ...
}
```

### Error Handling

Always handle errors explicitly. Never ignore them.

**Bad:**
```go
result, _ := doSomething() // Ignored error!
```

**Good:**
```go
result, err := doSomething()
if err != nil {
    log.Printf("failed to do something: %v", err)
    return fmt.Errorf("operation failed: %w", err)
}
```

### Context Usage

Use context for cancellation, deadlines, and request-scoped values.

```go
func processRequest(ctx context.Context, req Request) error {
    // Check for cancellation
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }

    // Pass context to downstream calls
    result, err := fetchData(ctx, req.ID)
    if err != nil {
        return err
    }

    return nil
}
```

### Struct Initialization

Use named fields for clarity.

**Bad:**
```go
todo := Todo{"1", "Task", "Description", false, time.Now()}
```

**Good:**
```go
todo := Todo{
    ID:          "1",
    Title:       "Task",
    Description: "Description",
    Completed:   false,
    CreatedAt:   time.Now(),
}
```

### Interface Design

Keep interfaces small and focused.

```go
// Good: Small, focused interfaces
type TodoStore interface {
    GetTodo(id string) (*Todo, error)
    SaveTodo(todo *Todo) error
}

type TodoValidator interface {
    Validate(todo *Todo) error
}
```

### Goroutine Management

Always ensure goroutines can terminate.

**Bad:**
```go
go func() {
    for {
        // Infinite loop with no way to stop!
        doWork()
        time.Sleep(time.Second)
    }
}()
```

**Good:**
```go
func startWorker(ctx context.Context) {
    go func() {
        ticker := time.NewTicker(time.Second)
        defer ticker.Stop()

        for {
            select {
            case <-ctx.Done():
                return // Clean exit
            case <-ticker.C:
                doWork()
            }
        }
    }()
}
```

## Frontend Best Practices

### JavaScript/Alpine.js

#### Component Structure

```javascript
// Good: Clear, organized Alpine component
<div x-data="{
    todos: [],
    loading: false,

    async loadTodos() {
        this.loading = true;
        try {
            this.todos = await fetchTodosFromDB();
        } catch (error) {
            console.error('Failed to load todos:', error);
        } finally {
            this.loading = false;
        }
    }
}" x-init="loadTodos()">
    <!-- Component template -->
</div>
```

#### Avoid Global State Pollution

```javascript
// Bad: Global variables
var currentUser = null;
var todos = [];

// Good: Encapsulated state
const todoApp = {
    state: {
        currentUser: null,
        todos: []
    },

    getTodos() {
        return this.state.todos;
    }
};
```

### HTML/HTMX

#### Semantic HTML

```html
<!-- Good: Semantic, accessible -->
<article class="todo-item" role="listitem">
    <h3 class="todo-title">Task Title</h3>
    <p class="todo-description">Description</p>
    <button class="btn-complete" aria-label="Mark as complete">
        Complete
    </button>
</article>
```

#### Progressive Enhancement

```html
<!-- Works without JavaScript, enhanced with it -->
<form action="/todos" method="POST"
      hx-post="/c/todo"
      hx-target="#todo-list"
      hx-swap="afterbegin">
    <input type="text" name="title" required>
    <button type="submit">Add Todo</button>
</form>
```

### CSS/Tailwind

#### Utility-First, Component-Second

```html
<!-- Good: Utility classes for one-offs -->
<div class="flex items-center justify-between p-4 bg-white rounded-lg shadow">
    <!-- Content -->
</div>

<!-- Good: Component classes for repeated patterns -->
<div class="todo-card">
    <!-- Content -->
</div>
```

```css
/* Define component classes in styles.css */
.todo-card {
    @apply flex items-center justify-between p-4 bg-white rounded-lg shadow;
}
```

## Code Organization

### File Structure

```
server/
├── server.go      # Server setup, initialization
├── routes.go      # HTTP handlers
├── types.go       # Data structures (if needed)
├── middleware/    # HTTP middleware
│   ├── auth.go
│   ├── logging.go
│   └── security.go
└── handlers/      # Complex handlers (if routes.go grows)
    ├── todo.go
    └── sync.go
```

### Function Length

Keep functions short and focused. Aim for < 50 lines.

**Bad:**
```go
func handleTodoRequest(w http.ResponseWriter, r *http.Request) {
    // 200 lines of code handling everything...
}
```

**Good:**
```go
func handleTodoRequest(w http.ResponseWriter, r *http.Request) {
    todo, err := parseTodoFromRequest(r)
    if err != nil {
        handleError(w, err)
        return
    }

    if err := validateTodo(todo); err != nil {
        handleValidationError(w, err)
        return
    }

    if err := saveTodo(todo); err != nil {
        handleSaveError(w, err)
        return
    }

    renderTodoResponse(w, todo)
}
```

### Cyclomatic Complexity

Reduce nested if statements and loops.

**Bad:**
```go
if condition1 {
    if condition2 {
        if condition3 {
            // Deeply nested!
            if condition4 {
                // Do something
            }
        }
    }
}
```

**Good:**
```go
if !condition1 {
    return
}
if !condition2 {
    return
}
if !condition3 {
    return
}
if !condition4 {
    return
}
// Do something
```

## Naming Conventions

### Variables

```go
// Bad
var d int                    // Too short, unclear
var numberOfTodosInDatabase int // Too verbose

// Good
var todoCount int
var userID string
var isCompleted bool
```

### Functions

Use verbs for actions, descriptive names for queries.

```go
// Actions
func saveTodo(todo *Todo) error
func deleteTodo(id string) error
func validateInput(data string) error

// Queries
func getTodoByID(id string) (*Todo, error)
func countActiveTodos() int
func hasPendingSync() bool
```

### Constants

```go
const (
    MaxTitleLength = 200
    DefaultTimeout = 30 * time.Second
    APIVersion     = "v1"
)
```

### Types

```go
// Structs: Nouns
type Todo struct { /* ... */ }
type TodoValidator struct { /* ... */ }

// Interfaces: Nouns or adjectives ending in "er"
type TodoStore interface { /* ... */ }
type Validator interface { /* ... */ }
type Serializable interface { /* ... */ }
```

## Error Handling

### Creating Errors

```go
import "errors"

// Simple errors
var ErrNotFound = errors.New("todo not found")
var ErrInvalidInput = errors.New("invalid input")

// Formatted errors
func validateTodo(todo *Todo) error {
    if todo.Title == "" {
        return fmt.Errorf("todo title is required")
    }
    return nil
}
```

### Wrapping Errors

```go
import "fmt"

func processTodo(id string) error {
    todo, err := fetchTodo(id)
    if err != nil {
        return fmt.Errorf("failed to fetch todo %s: %w", id, err)
    }

    // Process todo
    return nil
}
```

### Checking Error Types

```go
import "errors"

func handleError(err error) {
    if errors.Is(err, ErrNotFound) {
        // Handle not found
        return
    }

    var valErr *ValidationError
    if errors.As(err, &valErr) {
        // Handle validation error
        return
    }

    // Handle unknown error
}
```

## Testing Guidelines

### Test Structure

Follow the Arrange-Act-Assert pattern.

```go
func TestCreateTodo(t *testing.T) {
    // Arrange
    store := NewMemoryStore()
    todo := &Todo{
        ID:    "1",
        Title: "Test Todo",
    }

    // Act
    err := store.SaveTodo(todo)

    // Assert
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    retrieved, err := store.GetTodo("1")
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    if retrieved.Title != todo.Title {
        t.Errorf("expected title %q, got %q", todo.Title, retrieved.Title)
    }
}
```

### Table-Driven Tests

```go
func TestValidateTodoTitle(t *testing.T) {
    tests := []struct {
        name    string
        title   string
        wantErr bool
    }{
        {"valid title", "Valid Todo", false},
        {"empty title", "", true},
        {"too long", strings.Repeat("a", 201), true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateTodoTitle(tt.title)
            if (err != nil) != tt.wantErr {
                t.Errorf("validateTodoTitle() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Test Coverage

Aim for meaningful coverage, not just percentage.

```bash
# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Focus on testing:
- Critical business logic
- Error conditions
- Edge cases
- Public APIs

## Documentation Standards

### Package Documentation

```go
// Package server implements the HTTP server for Todo-Tasker.
// It provides route handlers, middleware, and template rendering.
package server
```

### Function Documentation

```go
// SaveTodo persists a todo to the store.
// It returns an error if the todo is invalid or if the save operation fails.
//
// Example:
//   todo := &Todo{ID: "1", Title: "Task"}
//   if err := store.SaveTodo(todo); err != nil {
//       log.Fatal(err)
//   }
func SaveTodo(todo *Todo) error {
    // Implementation
}
```

### Complex Logic

```go
// calculatePriority determines the priority score based on multiple factors:
// 1. Base priority set by user (0-3)
// 2. Days until due date (adds urgency)
// 3. Number of days overdue (increases priority)
func calculatePriority(todo *Todo) int {
    score := todo.Priority

    if todo.DueDate.Before(time.Now()) {
        // Overdue: add days overdue to score
        daysOverdue := int(time.Since(todo.DueDate).Hours() / 24)
        score += daysOverdue
    } else {
        // Not due: reduce score based on days until due
        daysUntilDue := int(time.Until(todo.DueDate).Hours() / 24)
        if daysUntilDue > 7 {
            score -= 1
        }
    }

    return score
}
```

## Git Workflow

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding/updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(server): add WebSocket sync endpoint

Implements the WebSocket endpoint for P2P device sync.
Includes session management and message routing.

Closes #42
```

```
fix(middleware): correct CORS header configuration

The Access-Control-Allow-Origin header was not being set
correctly for component requests.
```

### Branch Naming

```
feature/add-websocket-sync
fix/cors-headers
docs/update-installation-guide
refactor/simplify-routing
test/add-sync-tests
```

### Pull Request Process

1. Create feature branch from `main`
2. Make changes with clear commits
3. Write/update tests
4. Update documentation
5. Push and create PR
6. Address review feedback
7. Squash and merge

## Code Review Checklist

### Functionality
- [ ] Code does what it's supposed to do
- [ ] Edge cases are handled
- [ ] Error conditions are handled properly
- [ ] No obvious bugs

### Code Quality
- [ ] Code is readable and self-documenting
- [ ] Functions are small and focused
- [ ] No code duplication
- [ ] Follows project conventions
- [ ] No unnecessary complexity

### Testing
- [ ] Tests are included
- [ ] Tests cover happy path and error cases
- [ ] Tests are clear and maintainable
- [ ] All tests pass

### Security
- [ ] No SQL injection vulnerabilities
- [ ] No XSS vulnerabilities
- [ ] No command injection
- [ ] Input is validated
- [ ] Sensitive data is not logged

### Performance
- [ ] No obvious performance issues
- [ ] Database queries are efficient
- [ ] No memory leaks
- [ ] Goroutines can terminate properly

### Documentation
- [ ] Public functions are documented
- [ ] Complex logic is explained
- [ ] README is updated if needed
- [ ] Changelog is updated

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Clean Code by Robert C. Martin](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882)
- [The Pragmatic Programmer](https://pragprog.com/titles/tpp20/the-pragmatic-programmer-20th-anniversary-edition/)

## Getting Help

- Ask questions in pull requests
- Check [CONTRIBUTING.md](../CONTRIBUTING.md) for workflow
- Review existing code for examples
- Open issues for discussions

---

**[← Back to Get Started](get-started.md)**
