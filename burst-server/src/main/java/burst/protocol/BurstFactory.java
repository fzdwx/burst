package burst.protocol;

import com.google.protobuf.Any;
import com.google.protobuf.ByteString;
import com.google.protobuf.Int32Value;
import com.google.protobuf.StringValue;

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

    public static byte[] successForPort(final Map<Integer, Integer> portMap) {
        final var pack = Any.pack(Ports.newBuilder().putAllPorts(portMap).build());
        return BurstMessage.newBuilder()
                .setType(BurstType.INIT)
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
}