package burst.server.logic.trans;

import io.netty.channel.ChannelInboundHandlerAdapter;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 13:22
 */
public class TransformHandler extends ChannelInboundHandlerAdapter {

    private final String token;

    private final Integer availablePort;

    public TransformHandler(final String token, final Integer availablePort) {
        this.token = token;
        this.availablePort = availablePort;
    }
}