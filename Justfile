#!/usr/bin/env just --justfile

server:
    go run ./server/server.go -f=./server/etc/server.yaml

client:
    go run ./client/client.go -f=./client/etc/client.yaml -t dev

apis:
    goctl api go -api ./server/desc/server.api -dir ./server/

tidy:
    go mod tidy