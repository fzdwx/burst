package burst.inf.props;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

/**
 * burst config props.
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 12:02
 */
@Data
@ConfigurationProperties("server.burst")
public class BurstProps {


    /**
     * http port,启用表示所有http请求通过该端口访问,根据请求中的Host来区分；
     * 否则就是每个请求对应一个端口。(默认启用)
     */
    public int httpPort = 39399;
}