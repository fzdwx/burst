package burst.inf.metrics;

import lombok.Data;

import java.util.concurrent.atomic.LongAdder;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/23 21:41
 */
@Data
public class MemoryMetricsInfo implements MetricsInfo {

    /**
     * 读取字节数（从内网读取的）
     */
    private LongAdder readBytes;

    /**
     * 从内网读取了多少次
     */
    private LongAdder readCount;

    /**
     * 写字节数（写到内网的）
     */
    private LongAdder writeBytes;

    /**
     * 向内网写了多少次
     */
    private LongAdder writeCount;


}