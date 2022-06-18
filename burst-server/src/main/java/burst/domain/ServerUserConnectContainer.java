package burst.domain;

import core.Server;
import core.group.DefaultSocketGroup;
import core.group.SocketGroup;
import core.http.ext.WebSocket;
import core.socket.Socket;
import io.github.fzdwx.lambada.Collections;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Seq;
import io.github.fzdwx.lambada.anno.Nullable;
import io.netty.channel.Channel;
import io.netty.util.concurrent.GlobalEventExecutor;
import lombok.extern.slf4j.Slf4j;

import java.net.InetSocketAddress;
import java.util.Collection;
import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/25 21:37
 */
@Slf4j
public class ServerUserConnectContainer {

    private final static String ATTR_PORT = "port";
    private final SocketGroup<String> userConnectContainer = new DefaultSocketGroup<>(GlobalEventExecutor.INSTANCE);

    private final Map<ProxyInfo, Server> servers = Collections.map();
    /**
     * 与客户端的连接
     */
    private final WebSocket ws;
    private final String token;

    ServerUserConnectContainer(final WebSocket ws, final String token) {
        this.ws = ws;
        this.token = token;
    }

    public static ServerUserConnectContainer create(final WebSocket ws, final String token) {
        return new ServerUserConnectContainer(ws, token);
    }

    public String getToken(){
        return token;
    }

    public void addServer(final Map<ProxyInfo, Server> servers) {
        this.servers.putAll(servers);
    }

    @Nullable
    public Server getServer(final ProxyInfo proxyInfo) {
        return this.servers.get(proxyInfo);
    }

    public String addUserConnect(final Channel channel) {
        final var socket = Socket.create(channel);
        final var key = getKey(channel);
        socket.attr(ATTR_PORT, ((InetSocketAddress) channel.localAddress()).getPort());
        userConnectContainer.add(key, socket);

        log.debug("add user channel: {}", key);
        return key;
    }

    public boolean remove(final String userConnectId) {
        return this.userConnectContainer.remove(userConnectId);
    }

    /**
     * 返回与对应客户端的连接
     */
    public WebSocket ws() {
        return ws;
    }

    /**
     * 返回与对应客户端的连接（校验连通性）
     */
    public WebSocket safetyWs() {
        if (!this.ws.channel().isOpen() || !this.ws.channel().isActive()) {
            this.destroy();
            throw Exceptions.newIllegalState("客户端已经断开连接");
        }

        return ws;
    }

    /**
     * 回收资源,停止所有该客户端所被代理的关系
     */
    public void destroy() {
        closeServers(servers.values());
        userConnectContainer.close();
    }

    public Socket find(final String key) {
        return userConnectContainer.find(key);
    }

    public static void closeServers(Collection<Server> servers) {
        Seq.of(servers).nonNull().forEach(Server::close);
    }

    public void close(final Server server) {
        server.close().addListener(f -> {
            userConnectContainer.close(s -> s.attr(ATTR_PORT).equals(server.port()));
        });
    }

    private static String getKey(final Channel channel) {
        return channel.id().asLongText();
    }
}