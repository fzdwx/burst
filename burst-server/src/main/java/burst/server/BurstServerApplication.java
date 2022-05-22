package burst.server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.actuate.autoconfigure.metrics.web.servlet.WebMvcMetricsAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/15 21:43
 */
@SpringBootApplication(exclude = WebMvcMetricsAutoConfiguration.class)
public class BurstServerApplication {

    public static void main(String[] args) {
        final ConfigurableApplicationContext run = SpringApplication.run(BurstServerApplication.class);
    }

}