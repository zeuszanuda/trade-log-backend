FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/app-server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
COPY config ./config
COPY .env .

CMD ["./app"]
