# BUILD STAGE
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=dev

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=${VERSION}" \
    -o /app/service ./cmd/service

# FINAL STAGE
FROM alpine:latest

RUN adduser -D -H -s /sbin/nologin appuser

WORKDIR /

COPY --from=builder /app/service /service

USER appuser

EXPOSE 2112 6060

ENTRYPOINT ["/service"]
CMD ["-mode=fixed"]