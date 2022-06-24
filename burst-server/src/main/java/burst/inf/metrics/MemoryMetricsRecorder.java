package burst.inf.metrics;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * memory impl {@link MetricsRecorder }
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/23 21:25
 */
public class MemoryMetricsRecorder implements MetricsRecorder {

    private Map<String, MemoryMetricsInfo> cache;

    public MemoryMetricsRecorder() {
        cache = new ConcurrentHashMap<>();
    }

    @Override
    public MetricsInfo get(final String token) {
        return cache.computeIfAbsent(token, s -> new MemoryMetricsInfo());
    }

    @Override
    public void readFromClient(final String token, final String userConnectId, final int size) {
        get(token).incrRead(size);
    }

    @Override
    public void writeToUserError(final String token, final String userConnectId, final Throwable e) {
        get(token).incrRead(e);
    }

    @Override
    public void writeToClient(final String token, final String userConnectId, final int size) {
        get(token).incrWrite(size);
    }

    @Override
    public void writeToClientError(final String token, final String userConnectId, final Throwable e) {
        get(token).incrWrite(e);
    }
}