# [WIP] burst

Intranet to public network.

## Introduction
基本原理:
1. 客户端与服务端建立websocket连接并携带注册信息(需要暴露的端口)
2. 服务端随机监听(一个或多个,根据客户端的注册信息)端口,并把端口映射信息发送给客户端
3. 当有用户访问服务端的被映射的端口时,会转发到客户端
4. 客户端接收到请求后,会根据端口映射信息找到本地的端口(这个不一定是本机,后面可能会支持一个局域网内),
   然后与对应的端口进行连接,并发送请求,然后将请求转发给跟服务端的websocket连接


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
[click](https://github.com/fzdwx/burst/blob/main/Install.md)