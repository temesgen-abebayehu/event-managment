# Build stage
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# Runtime stage — use golang-alpine which already includes ca-certificates
FROM golang:1.23-alpine

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 3001
CMD ["./server"]
