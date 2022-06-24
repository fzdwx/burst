# 🚀 Quick Start

1. 准备一个服务端
    1. 当前提供了一个公有云的服务端: `114.132.249.192:10086 `
    2. ~~或者使用docker自建一个服务端: `docker run --name burst-server -p 10086:10086 -p 39399:39399 -d likelovec/burst-server:1.5`~~
       1. 这个不推荐，后来发现需要映射很多端口，很卡
2. [准备客户端](https://github.com/fzdwx/burst/releases)
3. 注册客户端，可以输入多个你想要被代理的端口信息，
   具体可以查看 [API 文档](https://www.apifox.cn/apidoc/shared-26c550f7-70a4-428b-8964-8f23c98b9abc/api-20962841)
4. 启动客户端 `./burst-client -sip 114.132.249.192 -sp 10086 -t {{token}}`
   ![image](https://user-images.githubusercontent.com/65269574/174085209-b9360ab9-bcd0-4e30-be0d-17018b058bc8.png)
5. 访问 `localhost:32988`

---

**如果还有不清楚的地方欢迎 [new issue](https://github.com/fzdwx/burst/issues/new)**