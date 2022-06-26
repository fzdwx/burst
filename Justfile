#!/usr/bin/env just --justfile

server:
    go run ./server/server.go -f=./server/etc/server.yaml

apis:
     goctl api go -api ./server/desc/server.api -dir ./server/

client:
    go run ./client/client.go -f=./client/etc/client.yaml

tidy:
    go mod tidy