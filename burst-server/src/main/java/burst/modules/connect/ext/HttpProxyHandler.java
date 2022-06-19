package burst.modules.connect.ext;

import burst.domain.ProxyInfo;
import burst.domain.ProxyType;
import burst.domain.ServerUserConnectContainer;
import burst.inf.props.BurstProps;
import burst.modules.connect.trans.HttpTransformHandler;
import burst.modules.connect.trans.Transform;
import core.Server;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/18 21:55
 */
public class HttpProxyHandler implements ProxyHandler {

    public HttpProxyHandler(final NioEventLoopGroup boss, final NioEventLoopGroup worker, final BurstProps burstProps) {
        startServer(boss, worker, burstProps).port();
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
        return null;
    }

    private Server startServer(final NioEventLoopGroup boss, final NioEventLoopGroup worker, final BurstProps burstProps) {
        final var server = new Server()
                .group(boss, worker)
                .childHandler(ch -> ch.pipeline().addLast(new ByteArrayDecoder(), new ByteArrayEncoder(), new HttpTransformHandler(burstProps)));
        server.listen(burstProps.http.port);
        return server;
    }
}