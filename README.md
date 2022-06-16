# :bug: **new issues** https://github.com/fzdwx/burst/issues/new

# burst

基于Java(Netty) & Go(net) 的内网穿透 or 反向代理程序。 [关于burst](https://github.com/fzdwx/burst/issues/6)

<br>

## ✨ Features

1. 支持所有基于tcp的协议 ( Support all tcp-based protocols )
2. 可以代理局域网(也可以理解成server可以访问的任意机器)内任意一台机器 (Proxy any machine in the LAN)
3. 动态添加以及关闭代理端口( Dynamically add and close proxy ports ) https://github.com/fzdwx/burst/issues/10#issuecomment-1153122850
5. ...

<br>

## 🚀 Quick Start
0. 准备服务端，当前提供了一个公有云的服务端addr: 114.132.249.192:10086
1. [下载客户端](https://github.com/fzdwx/burst/releases/tag/v1.0)
2. 注册，获取`token`(设置你想要被代理的机器的ip以及端口，可以输入多个)

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

3. 启动客户端

```shell
./burst-client -sip 114.132.249.192 -sp 10086 -t {{token}}
```

![image](https://user-images.githubusercontent.com/65269574/174085209-b9360ab9-bcd0-4e30-be0d-17018b058bc8.png)


_localhost:32988_ 就是最终代理到服务端的地址

<br>

## 👷 Install

[跳转](https://github.com/fzdwx/burst/blob/main/Install.md)
