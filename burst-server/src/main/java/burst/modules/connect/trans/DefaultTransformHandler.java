package burst.modules.connect.trans;

import burst.inf.metrics.MetricsRecorder;
import burst.protocol.BurstFactory;
import core.http.ext.WebSocket;
import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import lombok.extern.slf4j.Slf4j;

/**
 * 接收来自用户的请求,并转发。
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 13:22
 */
@Slf4j
public class DefaultTransformHandler extends BurstChannelHandler {

    private final Integer serverExportPort;
    private final WebSocket ws;
    private final String token;
    private final MetricsRecorder metricsRecorder;
    private String userConnectId;

    public DefaultTransformHandler(final Integer serverExportPort,
                                   final WebSocket webSocket,
                                   final String token,
                                   final MetricsRecorder metricsRecorder) {
        this.serverExportPort = serverExportPort;
        this.ws = webSocket;
        this.token = token;
        this.metricsRecorder = metricsRecorder;
    }

    @Override
    public void channelInactive(final ChannelHandlerContext ctx) throws Exception {
        Transform.remove(token, userConnectId);
        super.channelInactive(ctx);
    }

    @Override
    public void exceptionCaught(final ChannelHandlerContext ctx, final Throwable cause) throws Exception {
        Transform.remove(token, userConnectId);
        super.exceptionCaught(ctx, cause);
    }

    @Override
    protected void onUserConnect(final Channel channel) {
        userConnectId = Transform.add(channel, token);
        final var data = BurstFactory.userConnect(serverExportPort, userConnectId);

        ws.sendBinary(data);
    }

    @Override
    protected void onUserRequest(final Channel channel, final byte[] bytes) {
        final var data = BurstFactory.userRequest(userConnectId, bytes);
        ws.sendBinary(data).addListener(f -> {
            if (f.isSuccess()) {
                metricsRecorder.writeToClient(token, userConnectId, data.length);
            } else {
                metricsRecorder.writeToClientError(token, userConnectId, f.cause());
            }
        });
        log.info("user request size {}", bytes.length);
    }
}