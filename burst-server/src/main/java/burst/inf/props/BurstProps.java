package burst.inf.props;

import io.github.fzdwx.lambada.Collections;
import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

import java.util.Set;

/**
 * burst config props.
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 12:02
 */
@Data
@ConfigurationProperties(prefix = "server.burst")
public class BurstProps {

    public static BurstProps INS = new BurstProps();

    /**
     * tcp server worker cnt
     */
    public int workerCount = 0;

    /**
     * tcp server boss cnt
     */
    public int bossCount = 1;

    public Http http = new Http();

    @Data
    public static class Http {

        public boolean enable = true;
        /**
         * http port,启用表示所有http请求通过该端口访问,根据请求中的Host来区分；
         * 否则就是每个请求对应一个端口。(默认启用)
         */
        public int port = 39399;

        /**
         * 在header中host表示的字段，主要用来表示路由。(没有忽略大小写)
         */
        public Set<String> hostKeys = Collections.set("Host", ":authority");
    }
}