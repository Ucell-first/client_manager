FROM golang:1.24-alpine AS builder

WORKDIR /app

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/internal/delivery/http_cms/templates ./internal/delivery/http_cms/templates
COPY --from=builder /app/internal/delivery/http_cms/assets ./internal/delivery/http_cms/assets

CMD ["./main"]
