# BORK Frontend (TUI)

Terminal User Interface (TUI) for the BORK (Book Organization & Rental Kiosk) library management system.

## Technology Stack

- **Language**: Go 1.26.3
- **TUI Framework**: Bubble Tea (Elm-inspired)
- **Styling**: Lipgloss
- **HTTP Client**: Native Go HTTP client with cookie jar

## Project Structure

```
frontend/
├── cmd/
│   └── bork-tui/
│       └── main.go              # Application entry point
├── internal/
│   ├── ui/
│   │   ├── app.go               # Main application model
│   │   ├── login.go             # Login screen
│   │   └── ...                  # Future screens
│   ├── api/                     # API client (future)
│   └── config/                  # Configuration (future)
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── Dockerfile                   # Container image
└── README.md
```

## Prerequisites

- Go 1.26.3 or later
- Terminal with 256-color support
- Minimum terminal size: 80x24

## Running Locally

### 1. Install Dependencies

```bash
cd frontend
go mod download
```

### 2. Run Application

```bash
go run cmd/bork-tui/main.go
```

Or build and run:

```bash
go build -o bork-tui cmd/bork-tui/main.go
./bork-tui
```

## Building

```bash
# Build for current platform
go build -o bork-tui cmd/bork-tui/main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bork-tui-linux cmd/bork-tui/main.go

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o bork-tui-macos cmd/bork-tui/main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bork-tui.exe cmd/bork-tui/main.go

# Build Docker image
docker build -t bork-frontend:latest .
```

## Features (Current)

### Login Screen

- ASCII art logo
- Username input field
- Password input field (masked)
- Tab navigation between fields
- Visual focus indicators
- Keyboard shortcuts

### Keyboard Controls

- `Tab` / `Down Arrow` - Next field
- `Shift+Tab` / `Up Arrow` - Previous field
- `Enter` - Submit / Next field
- `Backspace` - Delete character
- `Esc` / `Ctrl+C` - Exit application

## Features (Future)

- [ ] API integration with backend
- [ ] Main dashboard
- [ ] Browse books screen
- [ ] Shopping cart
- [ ] Rental management
- [ ] Overdue notifications
- [ ] Configuration file support

## Development

### Adding a New Screen

1. Create a new file in `internal/ui/` (e.g., `dashboard.go`)
2. Implement the view struct with `Update()` and `View()` methods
3. Add the view type to `app.go`
4. Update routing in `app.go`

### Styling

Uses Lipgloss for styling. Common patterns:

```go
style := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FFFFFF")).
    Background(lipgloss.Color("#0000FF")).
    Bold(true).
    Padding(1, 2)
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestLoginView ./internal/ui
```

## Troubleshooting

### Terminal Display Issues

If the TUI doesn't display correctly:

```bash
# Check terminal capabilities
echo $TERM

# Set to 256-color mode
export TERM=xterm-256color

# Or
export TERM=screen-256color
```

### Build Issues

```bash
# Clean module cache
go clean -modcache

# Re-download dependencies
go mod download

# Verify dependencies
go mod verify
```

## Docker Usage

```bash
# Build image
docker build -t bork-frontend .

# Run interactively
docker run -it --rm bork-frontend

# Run with environment variables
docker run -it --rm -e API_BASE_URL=http://backend:8080/api bork-frontend
```

## Architecture

The application follows the Bubble Tea (Elm) architecture:

- **Model**: Application state (`App`, `LoginView`, etc.)
- **Update**: State transitions based on messages
- **View**: Render current state to string

```
┌─────────────┐
│   Message   │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Update    │ ──► New Model + Command
└─────────────┘
       │
       ▼
┌─────────────┐
│    View     │ ──► String (rendered UI)
└─────────────┘
```
