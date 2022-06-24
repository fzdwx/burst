package burst.inf.metrics;

import java.util.List;

/**
 * metrics info.
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/23 21:25
 */
public interface MetricsInfo {

    /**
     * 读取字节数（从内网读取的）
     */
    Long getReadBytes();

    /**
     * 从内网读取了多少次
     */
    Long getReadCount();

    /**
     * incr {@link #getReadBytes()} & {@link #getReadCount()}
     */
    void incrRead(int size);

    /**
     * 读取数据发生的错误（从内网读取的）
     */
    List<Throwable> getReadErrors();

    /**
     * 读取数据发生的错误次数
     */
    Long getReadErrorCount();

    void incrRead(Throwable e);

    /**
     * 写字节数（写到内网的）
     */
    Long getWriteBytes();

    /**
     * 向内网写了多少次
     */
    Long getWriteCount();

    /**
     * incr {@link #getWriteBytes()} & {@link #getWriteCount()}
     */
    void incrWrite(int size);

    /**
     * 写数据发生的错误（写到内网的）
     */
    List<Throwable> getWriteErrors();

    /**
     * 写数据发生的错误次数
     */
    Long getWriteErrorCount();

    void incrWrite(Throwable e);
}