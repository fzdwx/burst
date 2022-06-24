package burst.modules.connect.trans;

import burst.inf.metrics.MetricsRecorder;
import burst.inf.props.BurstProps;
import burst.protocol.BurstFactory;
import cn.hutool.core.io.IoUtil;
import io.github.fzdwx.lambada.Io;
import io.netty.channel.Channel;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import util.Netty;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.concurrent.atomic.AtomicIntegerFieldUpdater;

import static cn.hutool.core.io.IoUtil.lineIter;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 12:20
 */
@Slf4j
@RequiredArgsConstructor
public class HttpTransformHandler extends BurstChannelHandler {

    final static AtomicIntegerFieldUpdater<HttpTransformHandler> PARSED_STATE = AtomicIntegerFieldUpdater.newUpdater(HttpTransformHandler.class, "parsed");
    private final BurstProps burstProps;
    private final MetricsRecorder metricsRecorder;
    private String customDomain;
    private volatile int parsed = 0;
    private String userConnectId;


    @Override
    protected void onUserConnect(final Channel channel) {
    }

    @Override
    protected void onUserRequest(final Channel channel, final byte[] bytes) {
        if (initHost(bytes, channel)) {
            log.error("discard message: {}", channel.remoteAddress());
            return;
        }

        if (userConnectId == null) {
            notFound(channel);
            return;
        }

        final var container = Transform.getContainer(this.customDomain);
        if (container == null) {
            notFound(channel);
            return;
        }

        final var token = container.getToken();
        final var data = BurstFactory.userRequest(userConnectId, bytes);

        container.safetyWs().sendBinary(data)
                .addListener(f -> {
                    if (f.isSuccess()) {
                        metricsRecorder.writeToClient(token, userConnectId, bytes.length);
                    } else {
                        metricsRecorder.writeToClientError(token, userConnectId, f.cause());
                    }
                });
        log.info("user request size {}", bytes.length);
    }

    private void onUserConnect0(final Channel channel) {
        final var fakePort = Transform.getFakePort(customDomain);
        if (fakePort == null) {
            return;
        }

        final var container = Transform.getContainer(this.customDomain);
        if (container == null) {
            return;
        }

        final var ws = container.ws();
        if (!ws.channel().isOpen() || !ws.channel().isActive()) {
            container.destroy();
            Transform.removeCustomerContainerMapping(this.customDomain);
            Transform.destroy(container.getToken());
            return;
        }

        userConnectId = Transform.add(channel, container.getToken());
        final var data = BurstFactory.userConnect(fakePort, userConnectId);
        ws.sendBinary(data);
        log.info("user connect : customDomain={},fakePort={}", this.customDomain, fakePort);
    }

    // GET / HTTP/1.1
    // Host: qwe.com:9999
    // Connection: keep-alive
    // Cache-Control: max-age=0
    // Upgrade-Insecure-Requests: 1
    // User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36
    // Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
    // Accept-Encoding: gzip, deflate
    // Accept-Language: zh-CN,zh;q=0.9
    private boolean initHost(byte[] bytes, final Channel channel) {
        if (PARSED_STATE.compareAndSet(this, 0, 1)) {

            final BufferedReader reader = IoUtil.getReader(new InputStreamReader(Io.wrap(bytes)));

            for (String line : lineIter(reader)) {
                final String[] split = line.split(": ");
                if (split.length <= 1) {
                    continue;
                }

                if (burstProps.http.hostKeys.contains(split[0])) {
                    customDomain = split[1].split(":")[0];
                    onUserConnect0(channel);
                    break;
                }
            }
        }

        return customDomain == null;
    }

    private void notFound(final Channel channel) {
        final var resp = """
                HTTP/1.1 404 NOT FOUND\r
                transfer-encoding: chunked\r
                content-type: text/plain\r
                server: burst\r
                                
                0\r\n\r\n
                """;
        channel.writeAndFlush(Netty.wrap(channel.alloc(), resp));
    }
}