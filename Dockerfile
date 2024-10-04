# Build stage
FROM --platform=$BUILDPLATFORM golang:1.23-bookworm AS builder

# Install necessary build tools and Go
RUN apt-get update && apt-get install -y \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

ARG TARGETARCH
RUN GOOS=linux GOARCH=$TARGETARCH go build -o jbhelper cmd/main.go

RUN chmod +x jbhelper

# Clean up build artifacts
RUN go clean -cache -testcache

# Remove source code (keep only the compiled binary)
RUN rm -rf cmd internal

# Runtime stage
FROM debian:bookworm-slim AS runtime

WORKDIR /app

# Copy the compiled binary and static files from the builder stage
COPY --from=builder /app/jbhelper .
COPY --from=builder /app/static static
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the port
EXPOSE 10800

# Start the application
CMD ["/app/jbhelper", "-addr", ":10800"]
