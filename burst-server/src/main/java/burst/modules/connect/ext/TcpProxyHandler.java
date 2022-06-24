package burst.modules.connect.ext;

import burst.domain.ProxyInfo;
import burst.domain.ProxyType;
import burst.domain.ServerUserConnectContainer;
import burst.inf.metrics.MetricsRecorder;
import burst.modules.connect.trans.DefaultTransformHandler;
import core.Server;
import io.github.fzdwx.lambada.Exceptions;
import io.netty.channel.EventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import util.AvailablePort;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/18 21:32
 */
@Slf4j
@RequiredArgsConstructor
public class TcpProxyHandler implements ProxyHandler {

    private final EventLoopGroup boss;
    private final EventLoopGroup worker;

    private final MetricsRecorder metricsRecorder;

    @Override
    public String supportType() {
        return ProxyType.TCP;
    }

    @Override
    public Server apply(final String token, final ServerUserConnectContainer container, final ProxyInfo proxyInfo) {
        // todo 修改为配置化的
        final var availablePort = AvailablePort.random();
        if (availablePort == null) {
            log.error("[init] token={},host={}  port not available", token, proxyInfo);
            throw Exceptions.newIllegalState("服务端暂无可用端口");
        }
        final Server server = new Server()
                .group(boss, worker)
                .childHandler(ch -> ch.pipeline().addLast(
                        new ByteArrayDecoder(),
                        new ByteArrayEncoder(),
                        new DefaultTransformHandler(availablePort, container.safetyWs(), token, metricsRecorder)));
        server.listen(availablePort);
        proxyInfo.setServerExport(availablePort);
        log.info("client={},add {} proxy {} to {}", token, proxyInfo.type, proxyInfo.ip + ":" + proxyInfo.port, "localhost:" + server.port());
        return server;
    }
}