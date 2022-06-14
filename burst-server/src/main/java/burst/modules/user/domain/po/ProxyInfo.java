package burst.modules.user.domain.po;

import cn.hutool.core.text.StrPool;
import io.github.fzdwx.lambada.Assert;
import io.github.fzdwx.lambada.Exceptions;
import lombok.Data;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/12 13:36
 */
@Data
public class ProxyInfo {

    /**
     * 需要被代理的机器ip
     */
    public String ip = "localhost";

    /**
     * 对应端口号
     */
    public Integer port;

    public void preCheck() {
        Assert.notBlank(ip, "ip is required");

        if (port < 0 || port > 65535) {
            Exceptions.illegalArgument("port is not valid,must between 0 and 65535");
        }
    }

    @Override
    public String toString() {
        return ip + StrPool.COLON + port;
    }
}