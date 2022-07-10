# 快速开始

## 启动

1. 准备一个服务端和客户端 [下载](https://github.com/fzdwx/burst/releases)
2. 修改客户端的配置文件`client.yaml`,指定服务端的ip和端口
   ```yaml
   Server:
     # the server api port
     Port: 9999
     # the server ip
     Host: localhost
   ```
3. 启动服务端 `./server -l server.log`
    1. `-l` 指定客户端输出日志到文件,默认输出到控制台.
4. 启动客户端 `./client -l client.log -t xxxx`
    1. `-l` 指定客户端输出日志到文件,默认输出到控制台.
    2. `-t` 指定token，如果不指定，客户端会请求服务端生成token,token会是通过日志打印出来.后续所有API都需要携带这个token
       ```log
       2022/07/03 - 12:24:53 INF token: cb0ol5du3aotti323c8g
       ```

## 客户端CLI

> 现在启动客户端后,它就成了一个`CLI`,可以用来调用服务端的API.

比如说我需要添加代理: 我想暴露我本地的`63342`端口,你就可以输入`ap tcp::63342`,它就会发送请求到服务端,开启代理.

示例:

**add proxy**

![image](https://user-images.githubusercontent.com/65269574/178137099-de53f387-d321-4dfa-af41-0f9abeb426a5.png)

**show usage**

![image](https://user-images.githubusercontent.com/65269574/178137229-350ee6ff-382d-436e-bf44-9d325a780b7a.png)