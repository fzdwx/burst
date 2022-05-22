package burst.server.logic.trans;

import burst.protocol.BurstFactory;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import lombok.extern.slf4j.Slf4j;
import org.jetbrains.annotations.NotNull;
import socket.WebSocket;

/**
 * 接收来自用户的请求,并转发。
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 13:22
 */
@Slf4j
public class TransformHandler extends ChannelInboundHandlerAdapter {

    private final Integer serverExportPort;

    private final WebSocket ws;

    public TransformHandler(final Integer serverExportPort, final WebSocket webSocket) {
        this.serverExportPort = serverExportPort;
        this.ws = webSocket;
    }

    @Override
    public void channelActive(@NotNull final ChannelHandlerContext ctx) throws Exception {
        // step 2 [user connect] have user access, notify the client
        final var userConnectId = Transform.add(ctx.channel());
        final var data = BurstFactory.userConnect(serverExportPort, userConnectId);

        ws.sendBinary(data);

        log.info("user connect {}", ctx.channel().remoteAddress());
    }

    @Override
    public void channelRead(@NotNull final ChannelHandlerContext ctx, @NotNull final Object msg) throws Exception {
        // step 3 [user request] have user request, notify the client
        final var data = BurstFactory.userRequest(ctx.channel().id().asLongText(), (byte[]) msg);

        ws.sendBinary(data);

        log.info("user request {}", msg);
    }
}