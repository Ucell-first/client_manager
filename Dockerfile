FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/internal/delivery/http_cms/templates ./internal/delivery/http_cms/templates
COPY --from=builder /app/internal/delivery/http_cms/assets ./internal/delivery/http_cms/assets
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./myapp"]
