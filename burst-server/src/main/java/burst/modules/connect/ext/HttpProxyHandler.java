package burst.modules.connect.ext;

import burst.domain.ProxyInfo;
import burst.domain.ProxyType;
import burst.domain.ServerUserConnectContainer;
import burst.inf.props.BurstProps;
import burst.modules.connect.trans.HttpTransformHandler;
import burst.modules.connect.trans.Transform;
import core.Server;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;
import lombok.extern.slf4j.Slf4j;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/18 21:55
 */
@Slf4j
public class HttpProxyHandler implements ProxyHandler {

    private final int port;

    public HttpProxyHandler(final EventLoopGroup boss, final EventLoopGroup worker, final BurstProps burstProps) {
        this.port = startServer(boss, worker, burstProps).port();
    }

    @Override
    public String supportType() {
        return ProxyType.HTTP;
    }

    @Override
    public Server apply(final String token, final ServerUserConnectContainer container, final ProxyInfo proxyInfo) {
        Transform.saveCustomerMappingContainer(proxyInfo.customDomain, container);
        // fake
        proxyInfo.setServerExport(Transform.getFakePort(proxyInfo.customDomain));
        log.info("client={},add {} proxy {} to {}", token, proxyInfo.type, proxyInfo.ip + ":" + proxyInfo.port, "http://" + proxyInfo.customDomain + ":" + this.port);
        return null;
    }

    private Server startServer(final EventLoopGroup boss, final EventLoopGroup worker, final BurstProps burstProps) {
        final var server = new Server()
                .group(boss, worker)
                .childHandler(ch -> ch.pipeline().addLast(new ByteArrayDecoder(), new ByteArrayEncoder(), new HttpTransformHandler(burstProps)))
                .onSuccess(s -> {
                    log.info("http port start success");
                })
                .onFailure(f -> {
                    log.error("http port start failure", f);
                });
        server.listen(burstProps.http.port);
        Runtime.getRuntime().addShutdownHook(new Thread(server::shutdown));
        return server;
    }
}