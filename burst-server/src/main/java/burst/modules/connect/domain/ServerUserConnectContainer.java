package burst.modules.connect.domain;

import core.Server;
import core.group.DefaultSocketGroup;
import core.group.SocketGroup;
import core.http.ext.WebSocket;
import core.socket.Socket;
import io.github.fzdwx.lambada.Collections;
import io.netty.channel.Channel;
import io.netty.util.concurrent.GlobalEventExecutor;
import lombok.extern.slf4j.Slf4j;

import java.util.Collection;
import java.util.List;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/25 21:37
 */
@Slf4j
public class ServerUserConnectContainer {

    private final SocketGroup<String> userConnectContainer = new DefaultSocketGroup<>(GlobalEventExecutor.INSTANCE);

    private final List<Server> servers = Collections.list();
    /**
     * 与客户端的连接
     */
    private final WebSocket ws;

    ServerUserConnectContainer(final WebSocket ws) {
        this.ws = ws;
    }

    public static ServerUserConnectContainer create(final WebSocket ws) {
        return new ServerUserConnectContainer(ws);
    }

    public void addServer(final Collection<Server> servers) {
        this.servers.addAll(servers);
    }

    public String add(final Channel channel) {
        final var socket = Socket.create(channel);
        final var key = getKey(channel);
        userConnectContainer.add(key, socket);

        log.debug("add user channel: {}", key);
        return key;
    }

    /**
     * 返回与对应客户端的连接
     */
    public WebSocket ws() {
        return ws;
    }

    /**
     * 回收资源,停止所有该客户端所被代理的关系
     */
    public void destroy() {
        closeServers(servers);
        userConnectContainer.close();
    }

    public Socket find(final String key) {
        return userConnectContainer.find(key);
    }

    public static void closeServers(Collection<Server> servers) {
        for (final Server server : servers) {
            server.close();
        }
    }

    private static String getKey(final Channel channel) {
        return channel.id().asLongText();
    }
}