package burst.inf.config;

import burst.inf.props.BurstProps;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Configuration;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 12:03
 */
@Configuration
@EnableConfigurationProperties(BurstProps.class)
public class BurstAutoConfiguration {
}