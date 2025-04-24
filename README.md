# Store

A zero-cost database system implemented in Go, designed for simplicity and efficiency. It supports both local file storage and GitHub-based storage backends.

## Features

- 🚀 High-performance data storage and retrieval
- 💾 Multiple storage backends (Local, GitHub)
- 🔄 Data import/export capabilities
- 🔒 Built-in data validation and type checking
- 📊 Simple query and filtering support
- 🛠 Easy integration and usage
- 🔐 Secure authentication and authorization
- 📝 Comprehensive documentation

## Project Structure

```
store/
├── cmd/            # Application entry points
│   └── main.go     # Main application entry point
├── internal/       # Internal packages
│   ├── api/        # API handlers
│   ├── config/     # Configuration management
│   ├── handler/    # HTTP request handlers
│   ├── service/    # Business logic
│   ├── storage/    # Storage implementations
│   │   ├── local_store.go    # Local file storage
│   │   ├── github_store.go   # GitHub-based storage
│   │   └── store.go          # Storage interface
│   └── types/      # Data type definitions
├── examples/       # Usage examples
└── config.yaml     # Configuration file
```

## Quick Start

### Installation

```bash
go get github.com/Axpz/store
```

### Configuration

Create a `config.yaml` file:

```yaml
github:
  repo:
    owner: "your-github-username"
    name: "your-repo-name"
    branch: "main"
    tables:
      path: "tables"
      users: "users.json"
      comments: "comments.json"

server:
  port: 8080
  host: "localhost"

storage:
  type: "local" # or "github"
  path: "tables" # for local storage
```

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/Axpz/store/internal/config"
    "github.com/Axpz/store/internal/storage"
)

func main() {
    // Load configuration
    cfg, err := config.Load("config.yaml")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize store
    store, err := storage.New(cfg)
    if err != nil {
        log.Fatalf("Failed to create store: %v", err)
    }

    // Create a user
    user := storage.User{
        ID:       "user-1",
        Username: "john_doe",
        Email:    "john@example.com",
        Plan:     "free",
        Created:  time.Now().Unix(),
        Updated:  time.Now().Unix(),
    }
    
    err = store.Create(user)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }

    // Retrieve user
    retrieved, err := store.Get("user-1")
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }
    
    fmt.Printf("Retrieved user: %+v\n", retrieved)
}
```

### HTTP API

The project includes a RESTful HTTP API built with Gin:

```bash
# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username":"john_doe","email":"john@example.com","plan":"free"}'

# Get a user
curl http://localhost:8080/users/user-1

# Update a user
curl -X PUT http://localhost:8080/users/user-1 \
  -H "Content-Type: application/json" \
  -d '{"username":"john_doe","email":"john@example.com","plan":"premium"}'

# Delete a user
curl -X DELETE http://localhost:8080/users/user-1
```

## Development

### Requirements

- Go 1.21 or higher
- Supported operating systems: Linux, macOS, Windows
- For GitHub storage: GitHub API token

### Building

```bash
# Build all packages
make build

```

### Testing

```bash
# Run all tests
make test

# Run specific test
make test-pkg PKG=storage
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to all contributors who have helped shape this project
- Inspired by various database systems and storage solutions

