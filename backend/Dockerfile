# Multi-stage build for Spring Boot application

# Stage 1: Build
FROM maven:3.9-eclipse-temurin-25-alpine AS build
WORKDIR /app

# Copy pom.xml and download dependencies (cached layer)
COPY pom.xml .
RUN mvn dependency:go-offline -B

# Copy source code and checkstyle config, then build
COPY src ./src
COPY checkstyle.xml .

# Build with production profile, skip tests and checkstyle
RUN mvn clean package -DskipTests -Dcheckstyle.skip=true -Dmaven.test.skip=true \
    && mv target/*.jar target/app.jar

# Stage 2: Runtime
FROM eclipse-temurin:25-jre-alpine
WORKDIR /app

# Install dumb-init for proper signal handling
RUN apk add --no-cache dumb-init

# Create non-root user with specific UID/GID
RUN addgroup -g 1001 -S spring && adduser -u 1001 -S spring -G spring

# Copy JAR from build stage with specific name
COPY --from=build --chown=spring:spring /app/target/app.jar app.jar

# Switch to non-root user
USER spring:spring

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=40s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1

# Production JVM flags for optimization and security
ENV JAVA_OPTS="-XX:+UseContainerSupport \
    -XX:MaxRAMPercentage=75.0 \
    -XX:+UseG1GC \
    -XX:+UseStringDeduplication \
    -XX:+OptimizeStringConcat \
    -Djava.security.egd=file:/dev/./urandom \
    -Dspring.profiles.active=prod"

# Run application with dumb-init and optimized JVM settings
ENTRYPOINT ["dumb-init", "--"]
CMD ["sh", "-c", "java $JAVA_OPTS -jar app.jar"]
