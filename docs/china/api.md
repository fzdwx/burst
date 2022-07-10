# User

## 获取TOKEN

请求地址: `GET /user/auth`

响应:

```raw
cb0ojodu3aotti323c80
```

作用:

1. 用于启动客户端: `./client -t cb0ojodu3aotti323c80`
2. 后面所有的API请求都需要带上`TOKEN`
3. 如果启动时没有指定`TOKEN`,那么客户端将会自己请求服务端生成一个`TOKEN`,会通过日志打印出来
    ```raw
    2022/07/03 - 12:24:53 INF token: cb0ol5du3aotti323c8g
    ```
   用户后面如果调用`API`就要携带这个`TOKEN`

# Proxy

## 添加代理

请求地址: `POST /proxy/add/:token`

请求体(JSON):

```jsonpath
{
    "proxy": [
        {
            "ip":"localhost",    # 默认 为 localhost
            "port":63342,
            "channelType":"tcp", # 当前版本只支持 tcp
        }
    ]
}
```

响应:

```jsonpath
[
    {
        "ServerPort": 40477,                # 在服务端暴露的端口
        "IntranetAddr": "localhost:63342",  # 客户端中内网被代理的地址
        "ChannelType": "tcp"                    
    }
]
```

如果你代理的是`HTTP`服务,那么你现在你可以访问 `http://{{serverIp}}:40477`来检查是否代理成功.

## 删除代理

请求地址: `POST /proxy/remove/:token`

请求体(JSON):

```jsonpath
{
    "proxy": [
        {
            "ip":"localhost",    # 默认 为 localhost
            "port":63342,
            "channelType":"tcp", # 当前版本只支持 tcp
        }
    ]
}
```

响应:

```jsonpath
[
    {
        "ServerPort": 40477,                # 在服务端暴露的端口
        "IntranetAddr": "localhost:63342",  # 客户端中内网被代理的地址
        "ChannelType": "tcp"                    
    }
]
```

删除代理后，所有相关的连接都会立即关闭.