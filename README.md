# [WIP] burst

Intranet to public network.

## Introduction
todo


## Quick Start
1. [get client](https://github.com/fzdwx/burst/releases/tag/v1.0)
2. register your client
```shell
curl --location --request POST 'http://114.132.249.192:10086/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ports": [
        "63342"  // client wants exposed port
    ]
}'

// response
{"token":"348f952bb76e44d5a818440ef1bec53a"}
```
3. start your client
```shell
./burst-client -sip 114.132.249.192 -sp 10086 -t {{token}}
```
![img.png](img.png)

=> **114.132.249.192:28236** is your mapped address information

## Install
[click](https://github.com/fzdwx/burst/Install.md)