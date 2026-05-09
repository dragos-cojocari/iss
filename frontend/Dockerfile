# Multi-stage build for Go TUI application

# Stage 1: Build
FROM golang:1.23-alpine AS build
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bork-tui ./cmd/bork-tui

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from build stage
COPY --from=build /app/bork-tui .

# Run as non-root user
RUN addgroup -S bork && adduser -S bork -G bork
USER bork

# Set terminal environment
ENV TERM=xterm-256color

# Run the TUI application
ENTRYPOINT ["./bork-tui"]
