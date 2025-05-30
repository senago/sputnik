FROM golang:latest as builder

RUN apt-get update && \
    apt-get install -y gcc-mingw-w64-x86-64 && \
    rm -rf /var/lib/apt/lists/*

ENV GOOS=windows \
    GOARCH=amd64 \
    CGO_ENABLED=1 \
    CC=x86_64-w64-mingw32-gcc

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w -H=windowsgui" ./cmd/sputnik
