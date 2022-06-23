package burst.inf.metrics;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/23 21:25
 */
public class MemoryMetricsRecorder implements MetricsRecorder {

    @Override
    public void readFromClient(final String token, final String userConnectId, final int size) {

    }

    @Override
    public void writeToUserError(final String token, final String userConnectId, final Throwable e) {

    }


    @Override
    public void writeToClient(final String token, final String userConnectId, final int size) {

    }

    @Override
    public void writeToClientError(final String token, final String userConnectId, final Throwable e) {

    }
}