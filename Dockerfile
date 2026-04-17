# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bot ./cmd/bot

FROM alpine:3.22
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/bot .
VOLUME ["/app/data"]
ENV DATABASE_PATH=/app/data/registrations.db
ENTRYPOINT ["./bot"]
