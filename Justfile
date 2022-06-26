#!/usr/bin/env just --justfile

server:
    go run ./server/server.go -f=./server/etc/server.yaml

client:
    go run ./client/client.go -f=./client/etc/client.yaml -t dev

apis:
    goctl api go -api ./client/desc/client.api -dir ./client/

tidy:
    go mod tidy