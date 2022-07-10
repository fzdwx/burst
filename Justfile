#!/usr/bin/env just --justfile

amd64_linux := "GOOS=linux GOARCH=amd64"
amd64_win := "GOOS=windows GOARCH=amd64"

# list justfile commands
ls:
    @just -l

# install just to /usr/local/bin
prejust:
    cp -r just /usr/local/bin/just
    chmod +x /usr/local/bin/just

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
    cd ./cmd/server && {{amd64_linux}} go build -o server server.go  && cp server.example.yaml ../../bin/server.yaml && cp server ../../bin/server && cd ../../bin/ && tar -zvcf server-linux-amd64.tar.gz server server.yaml && rm -rf server && rm -rf server.yaml
    cd ./cmd/client && {{amd64_linux}} go build -o client client.go  && cp client.example.yaml ../../bin/client.yaml && cp client ../../bin/client && cd ../../bin/ && tar -zvcf client-linux-amd64.tar.gz client client.yaml && rm -rf client && rm -rf client.yaml
    cd ./cmd/server && {{amd64_win}} go build -o server.exe server.go  && cp server.example.yaml ../../bin/server.yaml && cp server.exe ../../bin/server.exe && cd ../../bin/ && tar -zvcf server-win-amd64.tar.gz server.exe server.yaml && rm -rf server.exe && rm -rf server.yaml
    cd ./cmd/client && {{amd64_win}} go build -o client.exe client.go  && cp client.example.yaml ../../bin/client.yaml && cp client.exe ../../bin/client.exe && cd ../../bin/ && tar -zvcf client-win-amd64.tar.gz client.exe client.yaml && rm -rf client.exe && rm -rf client.yaml
    cd ./cmd/client && rm -rf client && rm -rf client.exe
    cd ./cmd/server && rm -rf server && rm -rf server.exe
    @echo "Build success!"

build:
    cd ./cmd/server && {{amd64_linux}} go build -o ../../bin/server server.go
    cd ./cmd/client && {{amd64_linux}} go build -o ../../bin/client client.go

# call go mod tidy
tidy:
    go mod tidy