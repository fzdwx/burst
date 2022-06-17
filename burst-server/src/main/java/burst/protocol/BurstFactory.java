package burst.protocol;

import burst.domain.ProxyInfo;
import com.google.protobuf.Any;
import com.google.protobuf.ByteString;
import com.google.protobuf.Int32Value;
import com.google.protobuf.StringValue;
import io.github.fzdwx.lambada.Assert;
import io.github.fzdwx.lambada.Collections;

import java.util.List;
import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 13:33
 */
public class BurstFactory {


    public static byte[] error(BurstType type, String errorMessage) {
        return BurstMessage.newBuilder()
                .setType(type)
                .putHeader(Headers.ERROR.getNumber(), Any.pack(StringValue.of(errorMessage))).build().toByteArray();
    }

    public static byte[] successForPort(final Map<Integer, ProxyInfo> portMap) {
        final Map<Integer, Proxy> proxyMap = Collections.map();

        portMap.forEach((k, v) -> {
            proxyMap.put(k, Proxy.newBuilder().setPort(v.getPort()).setIp(v.getIp()).build());
        });


        final var pack = Any.pack(Ports.newBuilder().putAllPorts(proxyMap).build());
        return BurstMessage.newBuilder()
                .setType(BurstType.ADD_PROXY_INFO)
                .putHeader(Headers.PORTS.getNumber(), pack).build().toByteArray();
    }

    public static byte[] userConnect(final Integer serverExportPort, final String userConnectIdStr) {
        final var userConnectId = Any.pack(StringValue.of(userConnectIdStr));
        final var port = Any.pack(Int32Value.of(serverExportPort));
        return BurstMessage.newBuilder()
                .setType(BurstType.USER_CONNECT)
                .putHeader(Headers.SERVER_EXPORT_PORT.getNumber(), port)
                .putHeader(Headers.USER_CONNECT_ID.getNumber(), userConnectId)
                .build().toByteArray();
    }

    public static byte[] userRequest(final String userConnectIdStr, final byte[] data) {
        final var userConnectId = Any.pack(StringValue.of(userConnectIdStr));
        return BurstMessage.newBuilder()
                .setType(BurstType.FORWARD_DATA)
                .putHeader(Headers.USER_CONNECT_ID.getNumber(), userConnectId)
                .setData(ByteString.copyFrom(data)).build().toByteArray();
    }

    public static byte[] removeProxyInfo(final List<Integer> serverPorts) {
        Assert.notEmpty(serverPorts, "server ports is empty");
        return BurstMessage.newBuilder()
                .setType(BurstType.REMOVE_PROXY_INFO)
                .addAllServerPort(serverPorts).build().toByteArray();
    }
}