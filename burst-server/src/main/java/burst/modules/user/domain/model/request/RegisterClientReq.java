package burst.modules.user.domain.model.request;

import burst.inf.Info;
import cn.hutool.core.text.StrPool;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Lang;
import lombok.Data;

import java.util.Set;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 17:01
 */
@Data
public class RegisterClientReq implements Info {

    /**
     * 需要被代理的内网地址信息
     */
    private Set<Proxy> proxies;

    public void preCheck() {
        if (Lang.isEmpty(proxies)) {
            throw Exceptions.newIllegalArgument("port is empty");
        }

        for (final Proxy proxy : proxies) {
            proxy.preCheck();
        }

    }

    @Data
    public static class Proxy {

        /**
         * 需要被代理的机器ip
         */
        public String ip = "localhost";

        /**
         * 对应端口号
         */
        public Integer port;

        public void preCheck() {
            if (Lang.isBlank(ip)) {
                Exceptions.illegalArgument("ip is required");
            }

            if (port < 0 || port > 65535) {
                Exceptions.illegalArgument("port is not valid,must between 0 and 65535");
            }
        }

        @Override
        public String toString() {
            return ip + StrPool.COLON + port;
        }
    }
}