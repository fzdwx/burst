package burst.protocol;

import com.google.protobuf.Any;
import com.google.protobuf.StringValue;

import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 13:33
 */
public class BurstFactory {

    public static byte[] error(String errorMessage) {
        return BurstMessage.newBuilder().putHeader("error", Any.pack(StringValue.of(errorMessage))).build().toByteArray();
    }

    public static byte[] successForPort(final Map<Integer, Integer> portMap) {
        final var pack = Any.pack(Ports.newBuilder().putAllPorts(portMap).build());
        return BurstMessage.newBuilder().putHeader("ports", pack).build().toByteArray();
    }
}