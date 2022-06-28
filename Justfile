#!/usr/bin/env just --justfile

server:
    go run ./cmd/server.go

client:
    go run ./cmd/client.go

apis:
    goctl api go -api ./server/desc/server.api -dir ./server/

tidy:
    go mod tidy