# ─── Stage 1: Builder ───
FROM golang:1.25.1 AS builder

WORKDIR /app

# Optional: bypass TLS issues if behind a corporate proxy
ENV GOPROXY=direct
ENV GOSUMDB=off
ENV GOINSECURE=*
ENV GIT_SSL_NO_VERIFY=1

# Copy modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Expose port (documentation only, not actually published yet)
EXPOSE 8082

# Build a fully static binary (no GLIBC dependency)
# CGO_ENABLED=0 ensures pure Go static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/bin ./cmd/evaultz

# ─── Stage 2: Runtime ───
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy the static binary from builder
COPY --from=builder /app/bin /app/bin

# Expose port for documentation
EXPOSE 8082

# Run the app
ENTRYPOINT ["/app/bin"]
