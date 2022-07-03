# Quick start

## Prepare

1. Prepare a server and client [download](https://github.com/fzdwx/burst/releases)
2. Modify the client's configuration file `client.yaml`, specify the ip and port of the server

```yaml
Server:
  # the server api port
  Port: 9999
  # the server ip
  Host: localhost
```

## Add proxy 

request address: `POST /proxy/add/:token`

request body(JSON):

```jsonpath
{
    "proxy": [
        {
            "ip":"localhost",    # default is localhost
            "port":63342,
            "channelType":"tcp", # current version only supports tcp
        }
    ]
}
```

response:

```jsonpath
[
    {
        "ServerPort": 40477,                # 在服务端暴露的端口
        "IntranetAddr": "localhost:63342",  # 客户端中内网被代理的地址
        "ChannelType": "tcp"                    
    }
]
```

If you are proxying the `HTTP` service, then you can now visit `http://{{serverIp}}:40477` to check if the proxy is successful.