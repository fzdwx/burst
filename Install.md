## Install from release

[click](https://github.com/fzdwx/burst/releases/tag/v1.0)

## Installation from source code

1. clone source

```bash
git clone https://github.com/fzdwx/burst.git
```

2. package

```bash
cd burst && mvn -DskipTests clean package
```

3. run server

```bash
java -jar burst-server/target/burst-server-1.3.1.jar
```

4. register client

```shell
curl --location --request POST 'http://114.132.249.192:10086/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "proxies":[
        {
             "port":8080,
             "ip":"192.168.1.72"  # default is localhost
        }
    ]
}'

// response
{"token":"348f952bb76e44d5a818440ef1bec53a"}
```

5. run client

```shell
cd burst/burst-client
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