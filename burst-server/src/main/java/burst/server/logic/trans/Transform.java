package burst.server.logic.trans;

import burst.protocol.BurstMessage;
import burst.protocol.Headers;
import burst.server.logic.domain.model.ServerUserConnectContainer;
import burst.server.logic.domain.model.request.RegisterInfo;
import com.google.protobuf.StringValue;
import core.Server;
import core.socket.WebSocket;
import io.github.fzdwx.lambada.Collections;
import io.github.fzdwx.lambada.Lang;
import io.netty.channel.Channel;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import util.AvailablePort;

import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 12:44
 */
@Slf4j
public class Transform {

    private static final NioEventLoopGroup boss = new NioEventLoopGroup();
    private static final NioEventLoopGroup worker = new NioEventLoopGroup();
    private static final Map<String, ServerUserConnectContainer> serverContainer = Collections.map();

    /**
     * export ports.
     */
    public static Map<Integer, Integer> init(RegisterInfo info, WebSocket socket, String token) {
        final var container = ServerUserConnectContainer.create();
        final var portsMap = Collections.<Integer, Integer>map();

        for (Integer port : info.getPorts()) {
            final var availablePort = AvailablePort.random();
            if (availablePort == null) {
                log.error("[init] token={},port={}  port not available", token, port);
                return null;
            }

            final var server = new Server()
                    .group(boss, worker)
                    .childHandler(ch -> ch.pipeline().addLast(
                            new ByteArrayDecoder(),
                            new ByteArrayEncoder(),
                            new TransformHandler(availablePort, socket, token)
                    ));

            server.listen(availablePort);

            portsMap.put(availablePort, port);
            container.add(server);
        }

        serverContainer.put(token, container);
        log.info("client init ports:{}", portsMap);
        return portsMap;
    }

    /**
     * destroy server.
     *
     * @apiNote unbind port and close client channel.
     */
    public static void destroy(final String token) {
        final var container = serverContainer.remove(token);
        if (Lang.isNull(container)) {
            return;
        }
        container.destroy();
    }

    /**
     * when user connect, add to container.
     */
    public static String add(final Channel channel, final String token) {
        return serverContainer.get(token).add(channel);
    }

    @SneakyThrows
    public static void toUser(final BurstMessage burstMessage) {
        final var userConnectId = burstMessage.getHeaderMap().get(Headers.USER_CONNECT_ID.getNumber()).unpack(StringValue.class).getValue();
        final var token = burstMessage.getHeaderMap().get(Headers.TOKEN.getNumber()).unpack(StringValue.class).getValue();
        final var socket = serverContainer.get(token).find(userConnectId);

        if (socket == null) {
            log.debug("user not found:{}", userConnectId);
            return;
        }

        if (socket.channel().isActive()) {
            final var binary = burstMessage.getData().toByteArray();
            socket.send(binary);
            return;
        }

        log.debug("user not active:{}", userConnectId);
    }
}