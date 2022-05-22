package burst.server.logic.trans;

import burst.protocol.BurstMessage;
import burst.protocol.Headers;
import burst.server.logic.domain.model.request.RegisterInfo;
import com.google.protobuf.StringValue;
import core.Server;
import group.DefaultSocketGroup;
import group.SocketGroup;
import io.github.fzdwx.lambada.Collections;
import io.netty.channel.Channel;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;
import io.netty.util.concurrent.GlobalEventExecutor;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import socket.Socket;
import socket.WebSocket;
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

    private static final SocketGroup<String> userConnectContainer = new DefaultSocketGroup<>(GlobalEventExecutor.INSTANCE);

    /**
     * export ports.
     */
    public static Map<Integer, Integer> init(RegisterInfo info, WebSocket socket, String token) {
        Map<Integer, Integer> portsMap = Collections.map();

        for (Integer port : info.getPorts()) {

            final var availablePort = AvailablePort.random();
            if (availablePort == null) {
                return null;
            }

            portsMap.put(availablePort, port);

            // todo websocket 关闭后server也要关闭
            new Server().withGroup(boss, worker).withInitChannel(ch -> {
                ch.pipeline().addLast(new ByteArrayDecoder(), new ByteArrayEncoder(), new TransformHandler(availablePort, socket));
            }).listen(availablePort);
        }

        log.info("client init ports:{}", portsMap);

        return portsMap;
    }

    /**
     * when user connect, add to container.
     */
    public static String add(final Channel channel) {
        final var socket = Socket.create(channel);
        final var key = getKey(channel);
        userConnectContainer.add(key, socket);
        return key;
    }

    public static void remove(final Socket socket) {
        userConnectContainer.remove(getKey(socket.channel()));
    }

    @SneakyThrows
    public static void toUser(final BurstMessage burstMessage) {
        final var userConnectId = burstMessage.getHeaderMap().get(Headers.USER_CONNECT_ID.getNumber()).unpack(StringValue.class).getValue();
        final var socket = userConnectContainer.find(userConnectId);
        if (socket != null && socket.channel().isActive()) {
            final var binary = burstMessage.getData().toByteArray();
            socket.send(binary);
            return;
        }

        log.error("user not found:{}", userConnectId);
    }

    private static String getKey(final Channel channel) {
        return channel.id().asLongText();
    }
}