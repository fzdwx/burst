package burst.modules.connect.trans;

import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import lombok.extern.slf4j.Slf4j;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/17 16:43
 */
@Slf4j
public abstract class BurstChannelHandler extends ChannelInboundHandlerAdapter {

    protected abstract void onUserConnect(final Channel channel);

    protected abstract void onUserRequest(final Channel channel, final byte[] bytes);

    @Override
    public void channelActive(final ChannelHandlerContext ctx) throws Exception {
        final Channel channel = ctx.channel();
        onUserConnect(channel);
        log.info("user connect {}", channel.remoteAddress());
    }

    @Override
    public void channelRead(final ChannelHandlerContext ctx, final Object msg) throws Exception {
        onUserRequest(ctx.channel(), (byte[]) msg);
    }
}