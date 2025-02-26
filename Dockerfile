# Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/app .

EXPOSE 4000
CMD ["./app"]