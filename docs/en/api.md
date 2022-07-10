# User

## Get TOKEN

request address: `GET /user/auth`

response:

```raw
cb0ojodu3aotti323c80
```

usefulness:

1. used to start the client: `./client -t cb0ojodu3aotti323c80`
2. All subsequent API requests need to bring `TOKEN`
3. If `TOKEN` is not specified at startup, the client will request the server to generate a `TOKEN`, which will be
   printed out through the log
    ```raw
    2022/07/03 - 12:24:53 INF token: cb0ol5du3aotti323c8g
    ```
   If the user calls `API` later, they will carry this `TOKEN`

# Proxy

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

If you are proxying the `HTTP` service, then you can now visit `http://{{serverIp}}:40477` to check if the proxy is
successful.

## remove proxy

request address: `POST /proxy/remove/:token`

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

After the proxy is removed, all related connections are closed immediately.