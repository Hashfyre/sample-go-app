FROM golang:1.13.5-alpine3.11 AS builder

WORKDIR /app
COPY . .

RUN apk update \
    && apk add git \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-extldflags '-static'" -mod=vendor

FROM alpine:latest
COPY --from=builder /app/sample-go-app  /app/sample-go-app
WORKDIR /app

RUN apk update \
    && apk add --no-cache ca-certificates openssl \
    && update-ca-certificates
