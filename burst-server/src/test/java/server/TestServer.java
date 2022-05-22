package server;

import core.Server;
import io.netty.channel.nio.NioEventLoopGroup;
import org.junit.jupiter.api.Test;
import util.AvailablePort;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 12:14
 */
public class TestServer {

    @Test
    void name() {
        final var boss = new NioEventLoopGroup();
        final var worker = new NioEventLoopGroup();
        final var server = new Server()
                .withGroup(boss, worker)
                .onSuccess(s -> {
                    System.out.println(s.port());
                })
                .listen(0)
                .dispose();

    }
}