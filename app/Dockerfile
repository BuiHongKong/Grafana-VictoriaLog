# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY main.go .

RUN go mod init log-generator || true
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o log-generator main.go

# Run stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/log-generator .

CMD ["./log-generator"]
