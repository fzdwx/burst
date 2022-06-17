package burst.modules.connect.trans;

import burst.inf.props.BurstProps;
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
    private String host;
    private volatile int parsed = 0;

    public HttpTransformHandler(final String token, final BurstProps burstProps) {
        this.burstProps = burstProps;
    }

    @Override
    protected void onUserConnect(final Channel channel) {
        // todo 根据custom domain 获取websocket连接
    }

    @Override
    protected void onUserRequest(final Channel channel, final byte[] bytes) {
        if (initHost(bytes)) {
            log.error("discard message: {}", channel.remoteAddress());
        }
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

    public boolean initHost(byte[] bytes) {
        if (PARSED_STATE.compareAndSet(this, 0, 1)) {

            final BufferedReader reader = IoUtil.getReader(new InputStreamReader(Io.wrap(bytes)));

            for (String line : lineIter(reader)) {
                final String[] split = line.split(": ");
                if (split.length <= 1) {
                    continue;
                }

                if (burstProps.hostKeys.contains(split[0])) {
                    host = split[1];
                    onUserConnect(null);
                    break;
                }
            }
        }

        return host == null;
    }
}