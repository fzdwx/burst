#!/usr/bin/env just --justfile

amd64_linux := "GOOS=linux GOARCH=amd64"
amd64_win := "GOOS=windows GOARCH=amd64"

# list justfile commands
ls:
    @just -l

# run server
server:
    go run ./cmd/server.go

# run client
client:
    go run ./cmd/client.go

#build server and client binaries
build:
    cd ./cmd && {{amd64_linux}} go build -o ../bin/server-linux-amd64 ./server.go
    cd ./cmd && {{amd64_linux}} go build -o ../bin/client-linux-amd64 ./client.go
    cd ./cmd && {{amd64_win}} go build -o ../bin/server-win-amd64.exe ./server.go
    cd ./cmd && {{amd64_win}} go build -o ../bin/client-win-amd64.exe ./client.go

# call go mod tidy
tidy:
    go mod tidy