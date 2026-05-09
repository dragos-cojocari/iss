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
│   └── bork/
│       └── main.go              # Application entry point
├── internal/
│   ├── ui/
│   │   ├── app.go               # Main application model
│   │   ├── login.go             # Login screen
│   │   ├── overdue.go           # Overdue alert screen
│   │   ├── dashboard.go         # Main dashboard
│   │   ├── browse.go            # Browse books screen
│   │   ├── theme.go             # Matrix theme colors
│   │   └── quotes.go            # Quote manager
│   └── api/
│       ├── client.go            # HTTP client with cookies
│       ├── auth.go              # Authentication API
│       ├── books.go             # Books API
│       └── types.go             # API DTOs
├── quotes.txt                   # Matrix movie quotes
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
go run cmd/bork/main.go
```

Or build and run:

```bash
go build -o bork cmd/bork/main.go
./bork
```

## Building

```bash
# Build for current platform
go build -o bork cmd/bork/main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bork-linux cmd/bork/main.go

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o bork-macos cmd/bork/main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bork.exe cmd/bork/main.go

# Build Docker image
docker build -t bork-frontend:latest .
```

## Features

### Matrix Theme

- **Green-on-black aesthetic** inspired by The Matrix movie
- **Cycling quotes** from the movie (20 quotes total)
- **Neon highlights** for focused elements
- **Consistent styling** across all views

### Implemented Screens

#### 1. Login Screen

- Matrix-style ASCII logo with digital rain characters
- Username and password fields (masked)
- Tab navigation with visual focus indicators
- Enhanced button focus with neon border
- API integration with session cookies

#### 2. Overdue Alert Screen

- Displays overdue book notifications (currently dummy)
- Clear "All Clear" message when no overdue books
- Prominent Continue button

#### 3. Main Dashboard

- Current rentals display (0/3)
- Menu with 4 options:
  - Browse Books
  - My Rentals (future)
  - Return Books (future)
  - Logout
- Matrix quote that cycles on each visit
- Right-aligned username in header

#### 4. Browse Books Screen

- Paginated book list (4 books per page)
- Book details: title, author, category, ISBN, availability
- Navigation: ↑/↓ to select, ←/→ to change pages
- Quote cycles on page navigation
- Refresh functionality (R key)

### Keyboard Controls

**Login Screen:**

- `Tab` / `↓` - Next field
- `Shift+Tab` / `↑` - Previous field
- `Enter` - Submit / Next field
- `Esc` / `Ctrl+C` - Exit

**Dashboard:**

- `↑` / `↓` / `j` / `k` - Navigate menu
- `Enter` - Select menu item
- `Esc` - Logout and exit

**Browse Books:**

- `↑` / `↓` / `j` / `k` - Navigate books
- `←` / `→` / `h` / `l` / `PgUp` / `PgDn` - Change pages
- `R` - Refresh book list
- `B` / `Backspace` - Back to dashboard
- `Esc` - Back to dashboard

### Session Management

- Automatic logout on exit (Ctrl+C or Esc from dashboard)
- Session cookies managed transparently
- 30-minute session expiration

## Features (Future)

- [ ] My Rentals screen
- [ ] Return Books functionality
- [ ] Shopping cart for multiple rentals
- [ ] Search and filter books
- [ ] User profile management

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
