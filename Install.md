## Install from release

[click](https://github.com/fzdwx/burst/releases/tag/v1.0)

## Installation from source code

1. install protobuf

```bash
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
```

2. gen proto file

```shell
just protoc
```

3. build java code

- custom justFile set your mvn path
- add application-prod.yml

```shell
just build
just runServer
```

4. register client

```shell
curl --location --request POST 'http://114.132.249.192/:10086/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ports": [
        "63342"  // client wants exposed port
    ]
}'

// response
{"token":"348f952bb76e44d5a818440ef1bec53a"}
```

5. run client

```shell
cd burst-client
# get usage
go run . -u
go run . -sip {{serverIp}} -sp {{serverPort}} -t {{token}}
```

will output:

```json
{
  "level": "info",
  "message": "init success map[46233:63342]",
  "time": "2022-05-22 21:20:32 555"
}
```

46233 is the port exposed by server => **serverIp:46233**

to access `serverIp:46233` is to access `localhost:633242`