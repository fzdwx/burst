package burst.modules.connect.domain;

import core.Server;
import core.group.DefaultSocketGroup;
import core.group.SocketGroup;
import core.socket.Socket;
import io.github.fzdwx.lambada.Collections;
import io.netty.channel.Channel;
import io.netty.util.concurrent.GlobalEventExecutor;
import lombok.extern.slf4j.Slf4j;

import java.util.List;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/25 21:37
 */
@Slf4j
public class ServerUserConnectContainer {

    private final SocketGroup<String> userConnectContainer = new DefaultSocketGroup<>(GlobalEventExecutor.INSTANCE);

    private final List<Server> servers = Collections.list();

    public static ServerUserConnectContainer create() {
        return new ServerUserConnectContainer();
    }

    public void add(final Server server) {
        servers.add(server);
    }

    public String add(final Channel channel) {
        final var socket = Socket.create(channel);
        final var key = getKey(channel);
        userConnectContainer.add(key, socket);

        log.debug("add user channel: {}", key);
        return key;
    }

    public void destroy() {
        for (final Server server : servers) {
            server.close();
        }

        userConnectContainer.close();
    }

    public Socket find(final String key) {
        return userConnectContainer.find(key);
    }

    private static String getKey(final Channel channel) {
        return channel.id().asLongText();
    }
}