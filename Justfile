#!/usr/bin/env just --justfile

amd64_linux := "GOOS=linux GOARCH=amd64"
amd64_win := "GOOS=windows GOARCH=amd64"

# list justfile commands
ls:
    @just -l

# just run s -> just run server | just run c -> just run client
run type:
    ./just {{ if type == "s" { "server" } else { "client" } }}

# run server
server:
    cd ./cmd/server/ && go run .

# run client
client:
     cd ./cmd/client/ && go run .

#build server and client binaries
release:
    cd ./cmd/server && {{amd64_linux}} go build -o server server.go && tar -zvcf ../../bin/server-linux-amd64.tar.gz server server.yaml && rm -rf server
    cd ./cmd/client && {{amd64_linux}} go build -o client client.go && tar -zvcf ../../bin/client-linux-amd64.tar.gz client client.yaml && rm -rf client
    cd ./cmd/server && {{amd64_win}} go build -o server.exe server.go && tar -zvcf ../../bin/server-win-amd64.tar.gz server.exe server.yaml && rm -rf server.exe
    cd ./cmd/client && {{amd64_win}} go build -o client.exe client.go && tar -zvcf ../../bin/client-win-amd64.tar.gz client.exe client.yaml && rm -rf client.exe

build:
    cd ./cmd/server && {{amd64_linux}} go build -o ../../bin/server server.go
    cd ./cmd/client && {{amd64_linux}} go build -o ../../bin/client client.go

# call go mod tidy
tidy:
    go mod tidy