package burst.modules.connect.controller.trans;

import burst.protocol.BurstFactory;
import core.http.ext.WebSocket;
import io.github.fzdwx.lambada.anno.NotNull;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import lombok.extern.slf4j.Slf4j;

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
    private final String token;

    public TransformHandler(final Integer serverExportPort,
                            final WebSocket webSocket,
                            final String token) {
        this.serverExportPort = serverExportPort;
        this.ws = webSocket;
        this.token = token;
    }

    @Override
    public void channelActive(@NotNull final ChannelHandlerContext ctx) throws Exception {
        // step 2 [user connect] have user access, notify the client
        final var channel = ctx.channel();
        final var userConnectId = Transform.add(channel, token);
        final var data = BurstFactory.userConnect(serverExportPort, userConnectId);

        ws.sendBinary(data);

        log.info("user connect {}", channel.remoteAddress());
    }

    @Override
    public void channelRead(@NotNull final ChannelHandlerContext ctx, @NotNull final Object msg) throws Exception {
        // step 3 [user request] have user request, notify the client
        final var bytes = (byte[]) msg;
        final var data = BurstFactory.userRequest(ctx.channel().id().asLongText(), bytes);

        ws.sendBinary(data);
        log.info("user request size {}", bytes.length);
    }
}