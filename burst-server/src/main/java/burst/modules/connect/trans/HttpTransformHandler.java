package burst.modules.connect.trans;

import burst.inf.props.BurstProps;
import burst.protocol.BurstFactory;
import cn.hutool.core.io.IoUtil;
import io.github.fzdwx.lambada.Io;
import io.netty.channel.Channel;
import lombok.extern.slf4j.Slf4j;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.concurrent.atomic.AtomicIntegerFieldUpdater;

import static cn.hutool.core.io.IoUtil.lineIter;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 12:20
 */
@Slf4j
public class HttpTransformHandler extends BurstChannelHandler {

    final static AtomicIntegerFieldUpdater<HttpTransformHandler> PARSED_STATE = AtomicIntegerFieldUpdater.newUpdater(HttpTransformHandler.class, "parsed");
    private final BurstProps burstProps;
    private String customDomain;
    private volatile int parsed = 0;
    private String userConnectId;

    public HttpTransformHandler(final BurstProps burstProps) {
        this.burstProps = burstProps;
    }

    @Override
    protected void onUserConnect(final Channel channel) {
    }

    @Override
    protected void onUserRequest(final Channel channel, final byte[] bytes) {
        if (initHost(bytes, channel)) {
            log.error("discard message: {}", channel.remoteAddress());
            return;
        }

        final var data = BurstFactory.userRequest(userConnectId, bytes);
        Transform.getContainer(this.customDomain).safetyWs().sendBinary(data);
        log.info("user request size {}", bytes.length);
    }

    private void onUserConnect0(final Channel channel) {
        final var container = Transform.getContainer(this.customDomain);
        userConnectId = Transform.add(channel, container.getToken());
        final var fakePort = Transform.getFakePort(customDomain);
        final var data = BurstFactory.userConnect(fakePort, userConnectId);
        container.safetyWs().sendBinary(data);
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

                if (burstProps.hostKeys.contains(split[0])) {
                    customDomain = split[1].split(":")[0];
                    onUserConnect0(channel);
                    break;
                }
            }
        }

        return customDomain == null;
    }
}