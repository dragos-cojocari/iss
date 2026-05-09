# Build & Deployment Guide

This document describes the production-optimized Docker build process for the BORK application.

## Table of Contents

- [Overview](#overview)
- [Build Optimizations](#build-optimizations)
- [Backend Build](#backend-build)
- [Frontend Build](#frontend-build)
- [Building Images](#building-images)
- [Production Deployment](#production-deployment)
- [Troubleshooting](#troubleshooting)

## Overview

BORK uses multi-stage Docker builds to create minimal, secure, production-ready container images. The build process is optimized for:

- **Security**: Non-root users, stripped binaries, no debug symbols
- **Size**: Minimal runtime dependencies, no test or build artifacts
- **Performance**: Optimized JVM settings, static binaries
- **Reliability**: Proper signal handling, container-aware configurations

## Build Optimizations

### Common Optimizations

Both backend and frontend Dockerfiles implement:

1. **Multi-stage builds** - Separate build and runtime stages
2. **Layer caching** - Dependencies downloaded before source code copy
3. **Non-root users** - Specific UID/GID (1001) for security
4. **dumb-init** - Proper signal handling for graceful shutdowns
5. **Minimal base images** - Alpine Linux for small footprint
6. **No test artifacts** - Tests and linting skipped in production builds

## Backend Build

### Dockerfile Location

`backend/Dockerfile`

### Build Stage

**Base Image**: `maven:3.9-eclipse-temurin-25-alpine`

**Optimizations**:

- Dependencies downloaded separately for layer caching
- Tests skipped: `-DskipTests -Dcheckstyle.skip=true -Dmaven.test.skip=true`
- Checkstyle validation skipped (already done in CI/CD)
- Predictable JAR naming for deployment

```dockerfile
RUN mvn clean package -DskipTests -Dcheckstyle.skip=true -Dmaven.test.skip=true \
    && mv target/*.jar target/app.jar
```

### Runtime Stage

**Base Image**: `eclipse-temurin:25-jre-alpine`

**Key Features**:

1. **JVM Optimization Flags**:

   ```
   -XX:+UseContainerSupport       # Container-aware memory limits
   -XX:MaxRAMPercentage=75.0      # Use 75% of available RAM
   -XX:+UseG1GC                   # Modern garbage collector
   -XX:+UseStringDeduplication    # Reduce memory footprint
   -XX:+OptimizeStringConcat      # String optimization
   -Djava.security.egd=file:/dev/./urandom  # Faster startup
   -Dspring.profiles.active=prod  # Production profile
   ```

2. **Security**:

   - Non-root user `spring:spring` (UID/GID 1001)
   - Proper file ownership with `--chown`
   - No debug symbols or source code

3. **Signal Handling**:

   - `dumb-init` ensures proper SIGTERM/SIGINT handling
   - Enables graceful shutdowns in Kubernetes/Docker

4. **Health Check**:
   ```dockerfile
   HEALTHCHECK --interval=30s --timeout=3s --start-period=40s --retries=3 \
     CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1
   ```

### Image Size

- Build stage: ~500MB (discarded)
- Runtime image: ~200MB

## Frontend Build

### Dockerfile Location

`frontend/Dockerfile`

### Build Stage

**Base Image**: `golang:1.26.3-alpine`

**Optimizations**:

1. **Selective Source Copy** - Only production code:

   ```dockerfile
   COPY cmd ./cmd
   COPY internal ./internal
   # Tests and other files excluded
   ```

2. **Optimized Build Flags**:

   ```bash
   CGO_ENABLED=0              # No C dependencies
   GOOS=linux                 # Target Linux
   GOARCH=amd64              # Target AMD64
   -ldflags='-w -s -extldflags "-static"'  # Strip debug, static linking
   -trimpath                  # Remove file paths (security)
   ```

3. **Binary Size Reduction**:

   - `-w`: Strip DWARF debug info
   - `-s`: Strip symbol table
   - Result: ~30-50% smaller binary

4. **Dependency Verification**:
   ```dockerfile
   RUN go mod download && go mod verify
   ```

### Runtime Stage

**Base Image**: `alpine:3.19` (pinned version)

**Key Features**:

1. **Minimal Dependencies**:

   - `ca-certificates` - HTTPS support
   - `tzdata` - Timezone data
   - `dumb-init` - Signal handling
   - APK cache cleaned after install

2. **Security**:

   - Non-root user `bork:bork` (UID/GID 1001)
   - Static binary with no external dependencies
   - No debug symbols or file paths in binary

3. **Environment**:
   ```dockerfile
   ENV TERM=xterm-256color \
       TZ=UTC
   ```

### Image Size

- Build stage: ~400MB (discarded)
- Runtime image: ~15-20MB

## Building Images

### Prerequisites

- Docker 20.10+
- Docker Compose 2.0+
- Make

### Local Build

```bash
# Build all images
make docker-build

# Build specific service
docker compose build backend
docker compose build frontend
```

### Build Process

The `docker-build` target includes a workaround for Docker credential helper issues:

1. **Fetch base images** with temporary config:

   ```bash
   mkdir -p .docker-tmp && echo '{}' > .docker-tmp/config.json
   export DOCKER_CONFIG=.docker-tmp
   docker pull postgres:16-alpine
   docker pull maven:3.9-eclipse-temurin-25-alpine
   docker pull eclipse-temurin:25-jre-alpine
   docker pull golang:1.26.3-alpine
   docker pull alpine:3.19
   ```

2. **Build images** with Docker Compose:
   ```bash
   docker compose build
   ```

### CI/CD Build

The CI/CD pipeline (`cicd.yml`) uses **path-based filtering** so jobs only run when relevant files have changed. This is powered by [`dorny/paths-filter`](https://github.com/dorny/paths-filter).

#### Path Filters

| Filter     | Paths                                                                                       |
| ---------- | ------------------------------------------------------------------------------------------- |
| `backend`  | `backend/**`                                                                                |
| `frontend` | `frontend/**`                                                                               |
| `docker`   | `docker-compose.yml`, `docker-compose.dev.yml`, `backend/Dockerfile`, `frontend/Dockerfile` |

#### Conditional Job Execution

| Job             | Runs when                                                 |
| --------------- | --------------------------------------------------------- |
| `pre-commit`    | Always                                                    |
| `secret-scan`   | Always                                                    |
| `lint-backend`  | `backend/**` changed                                      |
| `lint-frontend` | `frontend/**` changed                                     |
| `test-backend`  | `backend/**` changed (after lint)                         |
| `test-frontend` | `frontend/**` changed (after lint)                        |
| `docker-build`  | Backend, frontend, or Docker config changed (after tests) |

Backend and frontend pipelines run **in parallel** — a PR touching only `backend/` will not wait for frontend lint or tests.

PRs that only modify documentation, database scripts, or other non-code files will skip lint, test, and docker-build jobs entirely.

#### Build & Publish (main branch)

The `build-publish.yml` workflow also uses path filters and only triggers on pushes to `main` that change `backend/**`, `frontend/**`, `docker-compose.yml`, or `docker-compose.dev.yml`.

#### Docker Build Step

```yaml
- name: Build Docker images
  run: make docker-build
```

The pipeline ensures:

- Code is linted (no Checkstyle in Docker build)
- Tests pass (no tests in Docker build)
- Secrets are scanned
- Only production-ready code is containerized

## Production Deployment

### Environment Variables

**Backend**:

```bash
SPRING_PROFILES_ACTIVE=prod
SPRING_DATASOURCE_URL=jdbc:postgresql://db:5432/bork
SPRING_DATASOURCE_USERNAME=bork
SPRING_DATASOURCE_PASSWORD=<secret>
```

**Frontend**:

```bash
BORK_BACKEND_URL=http://backend:8080
TERM=xterm-256color
TZ=UTC
```

### Resource Limits

**Backend** (recommended):

```yaml
resources:
  limits:
    memory: 512Mi
    cpu: 500m
  requests:
    memory: 256Mi
    cpu: 250m
```

**Frontend** (recommended):

```yaml
resources:
  limits:
    memory: 64Mi
    cpu: 100m
  requests:
    memory: 32Mi
    cpu: 50m
```

### Health Checks

**Backend**:

- Endpoint: `http://localhost:8080/api/health`
- Interval: 30s
- Timeout: 3s
- Start period: 40s
- Retries: 3

**Frontend**:

- No HTTP health check (TUI application)
- Process-based health monitoring recommended

### Security Considerations

1. **Run as non-root**: Both containers use UID/GID 1001
2. **Read-only filesystem**: Consider mounting `/app` as read-only
3. **Drop capabilities**: Use `--cap-drop=ALL` in production
4. **Network policies**: Restrict inter-service communication
5. **Secrets management**: Use Docker secrets or Kubernetes secrets

### Example Docker Compose (Production)

```yaml
services:
  backend:
    image: bork-backend:latest
    read_only: true
    cap_drop:
      - ALL
    security_opt:
      - no-new-privileges:true
    environment:
      SPRING_PROFILES_ACTIVE: prod
    tmpfs:
      - /tmp
    restart: unless-stopped

  frontend:
    image: bork-frontend:latest
    read_only: true
    cap_drop:
      - ALL
    security_opt:
      - no-new-privileges:true
    restart: unless-stopped
```

## Troubleshooting

### Docker Credential Helper Error

**Error**:

```
error getting credentials - err: exec: "docker-credential-desktop": executable file not found
```

**Solution**: Use the `docker-fetch-images` target which sets a temporary Docker config:

```bash
make docker-build  # Already includes docker-fetch-images
```

### Build Fails with "checkstyle.xml not found"

**Cause**: Checkstyle config not copied to build context

**Solution**: Ensure `COPY checkstyle.xml .` is in the Dockerfile before `RUN mvn package`

### Large Image Sizes

**Backend**:

- Expected: ~200MB
- If larger: Check for test artifacts or unnecessary dependencies

**Frontend**:

- Expected: ~15-20MB
- If larger: Ensure `-ldflags='-w -s'` is used in build

### Out of Memory Errors

**Backend**:

- Adjust `MaxRAMPercentage` in `JAVA_OPTS`
- Increase container memory limits
- Monitor with: `docker stats`

**Frontend**:

- Go binaries are typically memory-efficient
- Check for memory leaks in application code

### Slow Builds

**Solutions**:

1. Use layer caching - don't change `pom.xml` or `go.mod` frequently
2. Use BuildKit: `DOCKER_BUILDKIT=1 docker compose build`
3. Increase Docker resources in Docker Desktop settings
4. Use `--parallel` flag: `docker compose build --parallel`

### Permission Denied Errors

**Cause**: Files owned by root in container

**Solution**: Ensure `--chown=spring:spring` or `--chown=bork:bork` is used in COPY commands

## Best Practices

1. **Pin base image versions** - Use specific tags (e.g., `alpine:3.19` not `alpine:latest`)
2. **Scan for vulnerabilities** - Use `docker scan` or Snyk
3. **Multi-arch builds** - Consider ARM64 for Apple Silicon and AWS Graviton
4. **Image signing** - Use Docker Content Trust in production
5. **Registry security** - Use private registries with authentication
6. **Regular updates** - Rebuild images monthly for security patches

## Additional Resources

- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Spring Boot Docker Guide](https://spring.io/guides/topicals/spring-boot-docker/)
- [Go Docker Best Practices](https://docs.docker.com/language/golang/build-images/)
- [Alpine Linux Security](https://alpinelinux.org/about/)
