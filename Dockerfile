# Build stage
FROM --platform=$BUILDPLATFORM golang:1.23-bookworm AS builder

# Install necessary build tools and Go
RUN apt-get update && apt-get install -y \
    git \
    make \
    gcc \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

ARG TARGETARCH
RUN  make clean build

# Runtime stage
FROM debian:bookworm-slim AS runtime

WORKDIR /app

# Copy the compiled binary and static files from the builder stage
COPY --from=builder /app/bin/jbhelper .
COPY --from=builder /app/static static
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the port
EXPOSE 10800

VOLUME [ "/app/cert" ]

# Start the application
CMD ["/app/jbhelper", "-addr", ":10800"]
