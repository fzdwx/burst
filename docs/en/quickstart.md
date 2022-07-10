# Quick start

## start

1. Prepare a server and client [download](https://github.com/fzdwx/burst/releases)
2. Modify the client's configuration file `client.yaml`, specify the ip and port of the server
   ```yaml
   Server:
     # the server api port
     Port: 9999
     # the server ip
     Host: localhost
   ```
3. Start the server `./server -l server.log`
    1. `-l` specifies the log output file, defaults to the controlled output.
4. Start the client `./client -l client.log -t xxxx`
    1. `-l` specifies the log output file, defaults to the controlled output.
    2. `-t` specifies the token, if not specified, the client will request the server to generate a token, which will be
       printed out through the log,all subsequent APIs need to carry this token
       ```log
       2022/07/03 - 12:24:53 INF token: cb0ol5du3aotti323c8g
       ```

## Client CLI

> Now when the client is started, it becomes a `CLI` that can be used to call the server's API.

For example, I need to add a proxy: I want to expose my local `63342` port, you can enter `ap tcp::63342`, and it will
send a request to the server to enable the proxy.

Example:

**add proxy**

![image](https://user-images.githubusercontent.com/65269574/178137099-de53f387-d321-4dfa-af41-0f9abeb426a5.png)

**show usage**

![image](https://user-images.githubusercontent.com/65269574/178137229-350ee6ff-382d-436e-bf44-9d325a780b7a.png)