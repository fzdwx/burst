package burst.inf.metrics;

/**
 * 指标记录
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/23 21:17
 */
public interface MetricsRecorder {

    /**
     * response to user(read from client).
     * 在读取客户端数据时回调
     */
    void readFromClient(String token, String userConnectId, int size);

    /**
     * write to user error.
     */
    void writeToUserError(String token, String userConnectId, Throwable e);

    /**
     * user request(write to client).
     * 在写数据到客户端时回调.
     */
    void writeToClient(String token, String userConnectId, int size);

    /**
     * write to client error
     */
    void writeToClientError(String token, String userConnectId, Throwable e);
}