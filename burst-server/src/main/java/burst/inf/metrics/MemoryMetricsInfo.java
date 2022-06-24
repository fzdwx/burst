package burst.inf.metrics;

import lombok.Data;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.LongAdder;

/**
 * memory impl.
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/23 21:41
 */
@Data
public class MemoryMetricsInfo implements MetricsInfo {

    private LongAdder readBytes = new LongAdder();
    private LongAdder readCount = new LongAdder();
    private List<Throwable> readErrors = new ArrayList<>();
    private LongAdder readErrorCount = new LongAdder();
    private LongAdder writeBytes = new LongAdder();
    private LongAdder writeCount = new LongAdder();
    private List<Throwable> writeErrors = new ArrayList<>();
    private LongAdder writeErrorCount = new LongAdder();

    @Override
    public Long getReadBytes() {
        return this.readBytes.longValue();
    }

    @Override
    public Long getReadCount() {
        return this.readCount.longValue();
    }

    @Override
    public void incrRead(final int size) {
        this.readBytes.add(size);
        this.readCount.increment();
    }

    @Override
    public List<Throwable> getReadErrors() {
        return this.readErrors;
    }

    @Override
    public Long getReadErrorCount() {
        return this.readErrorCount.longValue();
    }

    @Override
    public void incrRead(final Throwable e) {
        this.readErrors.add(e);
        this.readErrorCount.increment();
    }

    @Override
    public Long getWriteBytes() {
        return this.writeBytes.longValue();
    }

    @Override
    public Long getWriteCount() {
        return this.writeCount.longValue();
    }

    @Override
    public void incrWrite(final int size) {
        this.writeBytes.add(size);
        this.writeCount.increment();
    }

    @Override
    public List<Throwable> getWriteErrors() {
        return this.writeErrors;
    }

    @Override
    public Long getWriteErrorCount() {
        return this.writeErrorCount.longValue();
    }

    @Override
    public void incrWrite(final Throwable e) {
        this.writeErrors.add(e);
        this.writeErrorCount.increment();
    }
}