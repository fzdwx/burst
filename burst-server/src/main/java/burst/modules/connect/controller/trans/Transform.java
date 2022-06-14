package burst.modules.connect.controller.trans;

import burst.modules.connect.domain.ServerUserConnectContainer;
import burst.modules.user.domain.model.request.RegisterClientReq;
import burst.modules.user.domain.po.ProxyInfo;
import burst.protocol.BurstFactory;
import burst.protocol.BurstMessage;
import burst.protocol.BurstType;
import burst.protocol.Headers;
import com.google.protobuf.StringValue;
import core.Server;
import core.http.ext.WebSocket;
import io.github.fzdwx.lambada.Collections;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Lang;
import io.github.fzdwx.lambada.Seq;
import io.netty.channel.Channel;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.handler.codec.bytes.ByteArrayDecoder;
import io.netty.handler.codec.bytes.ByteArrayEncoder;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import util.AvailablePort;

import java.util.Map;
import java.util.Set;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 12:44
 */
@Slf4j
public class Transform {

    private static final NioEventLoopGroup boss = new NioEventLoopGroup();
    private static final NioEventLoopGroup worker = new NioEventLoopGroup();
    private static final Map<String, ServerUserConnectContainer> serverContainer = Collections.map();

    /**
     * 初始化客户端的代理信息
     */
    public static void init(RegisterClientReq req, WebSocket ws, String token) {
        final var container = ServerUserConnectContainer.create(ws);
        final var older = serverContainer.put(token, container);

        addProxyInfo(token, req.getProxies());
        if (older != null) {
            older.destroy();
        }
    }

    /**
     * 添加代理信息,并发送消息到客户端
     *
     * @param token   token
     * @param proxies 代理信息
     * @apiNote 当该客户端断开连接或找不到可用端口时会返回null
     */
    public static void addProxyInfo(String token, Set<ProxyInfo> proxies) {
        if (Lang.isEmpty(proxies)) {
            return;
        }

        final var container = serverContainer.get(token);
        if (container == null) {
            return;
        }
        final var ws = container.safetyWs();

        final var portsMap = Collections.<Integer, ProxyInfo>map();
        final var servers = Collections.<ProxyInfo, Server>map();

        for (ProxyInfo proxyInfo : proxies) {
            final var availablePort = AvailablePort.random();
            if (availablePort == null) {
                log.error("[init] token={},host={}  port not available", token, proxyInfo);
                // 获取不到可用端口,回收当前监听的所有端口
                ServerUserConnectContainer.closeServers(servers.values());
                throw Exceptions.newIllegalState("服务端暂无可用端口");
            }

            final var server = new Server()
                    .group(boss, worker)
                    .childHandler(ch -> ch.pipeline().addLast(
                            new ByteArrayDecoder(),
                            new ByteArrayEncoder(),
                            new TransformHandler(availablePort, ws, token)
                    ));

            server.listen(availablePort);
            servers.put(proxyInfo, server);
            portsMap.put(availablePort, proxyInfo);
        }
        container.addServer(servers);

        if (portsMap.isEmpty()) {
            ws.sendBinary(BurstFactory.error(BurstType.ADD_PROXY_INFO, "portMap is null,maybe server did not have available Port"));
            return;
        }

        log.info("client add proxy ports:{}", portsMap);
        ws.sendBinary(BurstFactory.successForPort(portsMap));
    }

    /**
     * 移除代理信息,并发送消息到客户端
     *
     * @param token   token
     * @param proxies 代理信息
     */
    public static void removeProxyInfo(final String token, final Set<ProxyInfo> proxies) {
        if (Lang.isEmpty(proxies)) {
            return;
        }

        final var container = serverContainer.get(token);
        if (container == null) {
            return;
        }

        final var ws = container.safetyWs();

        final var serverPorts = Seq.of(proxies)
                .map(container::getServer)
                .nonNull()
                .and(container::close)
                .map(Server::port)
                .toList();

        if (serverPorts.isEmpty()) {
            throw Exceptions.newIllegalState("没有需要关闭的服务端端口映射!");
        }

        // notify client remove proxy info
        ws.sendBinary(BurstFactory.removeProxyInfo(serverPorts));
    }

    /**
     * destroy server.
     *
     * @apiNote unbind port and close client channel.
     */
    public static void destroy(final String token) {
        final var container = serverContainer.remove(token);
        if (Lang.isNull(container)) {
            return;
        }
        container.destroy();
    }

    /**
     * bind user channel to client connect container.
     */
    public static String add(final Channel channel, final String token) {
        return serverContainer.get(token).addUserConnect(channel);
    }

    @SneakyThrows
    public static void toUser(final BurstMessage burstMessage) {
        final var userConnectId = burstMessage.getHeaderMap().get(Headers.USER_CONNECT_ID.getNumber()).unpack(StringValue.class).getValue();
        final var token = burstMessage.getHeaderMap().get(Headers.TOKEN.getNumber()).unpack(StringValue.class).getValue();
        final var socket = serverContainer.get(token).find(userConnectId);

        if (socket == null) {
            log.debug("user not found:{}", userConnectId);
            return;
        }

        if (socket.channel().isActive()) {
            final var binary = burstMessage.getData().toByteArray();
            socket.send(binary);
            return;
        }

        log.debug("user not active:{}", userConnectId);
    }
}