# 快速开始

## 准备

1. 准备一个服务端和客户端 [下载](https://github.com/fzdwx/burst/releases)
2. 修改客户端的配置文件`client.yaml`,指定服务端的ip和端口

```yaml
Server:
  # the server api port
  Port: 9999
  # the server ip
  Host: localhost
```

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