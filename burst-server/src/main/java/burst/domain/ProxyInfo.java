package burst.domain;

import burst.modules.connect.trans.Transform;
import cn.hutool.core.text.StrPool;
import io.github.fzdwx.lambada.Assert;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Lang;
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

    /**
     * 当前通道的类型
     *
     * @see ProxyType
     */
    public String type = ProxyType.TCP;

    /**
     * 自定义域,当{@link #type} 为 {@link ProxyType#HTTP} 时才有效。
     */
    public String customDomain = "";

    /**
     * 服务器暴露的端口
     */
    private Integer serverExport;

    public void preCheck() {
        Assert.notBlank(ip, "ip is required");
        Assert.notBlank(this.type, "the type is required");

        if (port < 0 || port > 65535) {
            Exceptions.illegalArgument("port is not valid,must between 0 and 65535");
        }

        if (Lang.eq(this.type, ProxyType.HTTP)) {
            if (Lang.isBlank(this.customDomain)) {
                Exceptions.illegalArgument("type is http,must set customDomain");
            }

            if (Transform.hasCustomDomain(this.customDomain)) {
                Exceptions.illegalArgument("the customDomain is duplicated");
            }

            Transform.putCustomDomain(this.customDomain);
        }
    }

    @Override
    public String toString() {
        return ip + StrPool.COLON + port;
    }

    public void setServerExport(final Integer availablePort) {
        this.serverExport = availablePort;
    }
}