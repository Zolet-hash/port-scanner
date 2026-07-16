# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Download dependencies first (better caching)
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Runtime stage (small image)
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]