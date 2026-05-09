# Multi-stage build for Go TUI application

# Stage 1: Build
FROM golang:1.26.3-alpine AS build
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy go mod files and download dependencies (cached layer)
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code (exclude tests and unnecessary files)
COPY cmd ./cmd
COPY internal ./internal

# Build optimized binary for production
# -ldflags: strip debug info and reduce binary size
# -trimpath: remove file system paths from binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a \
    -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -trimpath \
    -o bork-tui ./cmd/bork-tui

# Stage 2: Runtime (minimal scratch-based image)
FROM alpine:3.19
WORKDIR /app

# Install minimal runtime dependencies
RUN apk add --no-cache ca-certificates tzdata dumb-init \
    && rm -rf /var/cache/apk/*

# Create non-root user with specific UID/GID
RUN addgroup -g 1001 -S bork && adduser -u 1001 -S bork -G bork

# Copy binary from build stage
COPY --from=build --chown=bork:bork /app/bork-tui .

# Switch to non-root user
USER bork:bork

# Set production environment variables
ENV TERM=xterm-256color \
    TZ=UTC

# Run the TUI application with dumb-init for proper signal handling
ENTRYPOINT ["dumb-init", "--"]
CMD ["./bork-tui"]
