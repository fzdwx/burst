package burst.inf.config;

import burst.inf.metrics.MemoryMetricsRecorder;
import burst.inf.metrics.MetricsRecorder;
import burst.inf.props.BurstProps;
import burst.modules.connect.ext.HttpProxyHandler;
import burst.modules.connect.ext.ProxyHandler;
import burst.modules.connect.ext.TcpProxyHandler;
import burst.modules.connect.trans.Transform;
import io.github.fzdwx.lambada.Collections;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.epoll.Epoll;
import io.netty.channel.epoll.EpollEventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.List;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 12:03
 */
@Configuration
@EnableConfigurationProperties(BurstProps.class)
public class BurstAutoConfiguration {

    @Bean(destroyMethod = "shutdownGracefully")
    public EventLoopGroup boss(BurstProps burstProps) {
        if (Epoll.isAvailable()) {
            return new EpollEventLoopGroup(burstProps.bossCount);
        }
        return new NioEventLoopGroup(burstProps.bossCount);
    }

    @Bean(destroyMethod = "shutdownGracefully")
    public EventLoopGroup worker(BurstProps burstProps) {
        if (Epoll.isAvailable()) {
            return new EpollEventLoopGroup(burstProps.workerCount);
        }
        return new NioEventLoopGroup(burstProps.workerCount);
    }

    /**
     * @see Transform#setApplicationContext(ApplicationContext)
     */
    @Bean
    public List<ProxyHandler> proxyHandlers(EventLoopGroup boss,
                                            EventLoopGroup worker,
                                            BurstProps burstProps,
                                            MetricsRecorder metricsRecorder) {
        final var tcpProxyHandler = new TcpProxyHandler(boss, worker, metricsRecorder);
        final var list = Collections.<ProxyHandler>list(tcpProxyHandler);
        if (burstProps.http.enable) {
            list.add(new HttpProxyHandler(boss, worker, burstProps, metricsRecorder));
        }
        return list;
    }

    @Bean
    public Transform transform() {
        return new Transform();
    }

    @Bean
    public MemoryMetricsRecorder metricsRecorder() {
        return new MemoryMetricsRecorder();
    }
}