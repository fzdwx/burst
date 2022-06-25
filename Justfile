#!/usr/bin/env just --justfile

run:
    go run burst.go -f ./etc/burst-api.yaml

tidy:
    go mod tidy