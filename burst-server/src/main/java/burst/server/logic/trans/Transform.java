package burst.server.logic.trans;

import burst.server.logic.domain.model.request.RegisterInfo;
import core.Server;
import group.DefaultSocketGroup;
import group.SocketGroup;
import io.github.fzdwx.lambada.Collections;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;
import io.netty.util.concurrent.GlobalEventExecutor;
import lombok.extern.slf4j.Slf4j;
import socket.WebSocket;
import util.AvailablePort;

import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 12:44
 */
@Slf4j
public class Transform {

    private static final NioEventLoopGroup boss = new NioEventLoopGroup(1);
    private static final NioEventLoopGroup worker = new NioEventLoopGroup();

    private static final Map<String, Map<Integer, Integer>> clientPortMapping = Collections.map();
    private static final SocketGroup<String> group = new DefaultSocketGroup<>(GlobalEventExecutor.INSTANCE);

    public static Map<Integer, Integer> init(RegisterInfo info, WebSocket socket, String token) {
        Map<Integer, Integer> portsMap = Collections.map();
        for (Integer port : info.getPorts()) {

            final var availablePort = AvailablePort.random();
            if (availablePort == null) {
                return null;
            }

            portsMap.put(availablePort, port);

            new Server()
                    .withGroup(boss, worker)
                    .withInitChannel(ch -> {
                        ch.pipeline().addLast(
                                new ByteArrayDecoder(),
                                new ByteArrayEncoder(),
                                new TransformHandler(token, availablePort)
                        );
                    })
                    .listen(availablePort);
        }
        group.add(token, socket);

        log.info("client init ports:{}", portsMap);

        return portsMap;
    }

    public static void remove(final String token) {
        group.remove(token);
    }
}