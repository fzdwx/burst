package burst.modules.connect.trans;

import burst.domain.ProxyInfo;
import burst.domain.ProxyType;
import burst.domain.ServerUserConnectContainer;
import burst.domain.model.request.RegisterClientReq;
import burst.modules.connect.ext.ProxyHandler;
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
import io.github.fzdwx.lambada.anno.Nullable;
import io.netty.channel.Channel;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.BeansException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;

import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.atomic.LongAdder;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 12:44
 */
@Slf4j
public class Transform implements ApplicationContextAware {

    private static final Map<String, ServerUserConnectContainer> serverContainer = Collections.map();
    private static final Map<String, ServerUserConnectContainer> customDomainMappingClient = Collections.map();
    /**
     * port is fake
     */
    private static final Map<String, Integer> customDomainMappingPort = Collections.map();
    private static final LongAdder ADDER = new LongAdder();
    private static Map<String, ProxyHandler> proxyHandlers;

    public static Integer putCustomDomain(String domain) {
        ADDER.decrement();
        return customDomainMappingPort.put(domain, ADDER.intValue());
    }

    @Nullable
    public static Integer getFakePort(String customDomain) {
        return customDomainMappingPort.get(customDomain);
    }

    public static Integer removeFakePort(String customDomain) {
        return customDomainMappingPort.remove(customDomain);
    }

    public static boolean hasCustomDomain(String customDomain) {
        return customDomainMappingClient.containsKey(customDomain);
    }

    public static ServerUserConnectContainer getContainer(String customDomain) {
        return customDomainMappingClient.get(customDomain);
    }

    public static ServerUserConnectContainer removeCustomerContainerMapping(String customDomain) {
        return customDomainMappingClient.remove(customDomain);
    }

    public static ServerUserConnectContainer saveCustomerMappingContainer(String customDomain, ServerUserConnectContainer container) {
        return customDomainMappingClient.put(customDomain, container);
    }

    /**
     * 初始化客户端的代理信息
     */
    public static void init(RegisterClientReq req, WebSocket ws, String token) {
        final var container = ServerUserConnectContainer.create(ws, token);
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

        final WebSocket ws = container.safetyWs();
        final Map<Integer, ProxyInfo> portsMap = listenAndGetPortMapping(container, token, proxies);

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

        final var serverPorts = Seq.of(proxies).map(proxyInfo -> {
            final var server = container.getServer(proxyInfo);
            if (server != null) {
                server.close();
            }

            if (proxyInfo.getType().equals(ProxyType.HTTP)) {
                final var fakePort = getFakePort(proxyInfo.customDomain);
                removeFakePort(proxyInfo.customDomain);
                removeCustomerContainerMapping(proxyInfo.customDomain);
                return fakePort;
            }

            // not null.
            return server.port();
        }).toList();

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

    /**
     * remove user connect from container.
     */
    public static boolean remove(String token, String userConnectId) {
        final ServerUserConnectContainer container = serverContainer.get(token);
        if (container == null) {
            return false;
        }

        return container.remove(userConnectId);
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

    @Override
    public void setApplicationContext(final ApplicationContext applicationContext) throws BeansException {
        proxyHandlers = Seq.toMap((List<ProxyHandler>) applicationContext.getBean("proxyHandlers"), ProxyHandler::supportType);
    }

    private static Map<Integer, ProxyInfo> listenAndGetPortMapping(ServerUserConnectContainer container, final String token,
                                                                   final Set<ProxyInfo> proxies) {
        final var ws = container.safetyWs();
        final var portsMap = Collections.<Integer, ProxyInfo>map();
        final var servers = Collections.<ProxyInfo, Server>map();
        try {
            for (ProxyInfo proxyInfo : proxies) {
                final var proxyHandler = proxyHandlers.get(proxyInfo.type);
                if (proxyHandler == null) {
                    Exceptions.illegalArgument("unSupport type:" + proxyInfo.type);
                }

                final var server = proxyHandler.apply(token, container, proxyInfo);
                if (server != null) {
                    servers.put(proxyInfo, server);
                }

                portsMap.put(proxyInfo.getServerExport(), proxyInfo);
            }
        } catch (Exception e) {
            ServerUserConnectContainer.closeServers(servers.values());
            log.error("get server error", e);
        }

        container.addServer(servers);

        return portsMap;
    }
}