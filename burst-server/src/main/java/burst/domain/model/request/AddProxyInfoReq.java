package burst.domain.model.request;

import burst.domain.ProxyInfo;
import burst.temp.Cache;
import io.github.fzdwx.lambada.Assert;
import lombok.Data;

import java.util.Set;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/12 13:56
 */
@Data
public class AddProxyInfoReq {

    /**
     * 令牌
     */
    private String token;

    /**
     * 想要添加的代理信息
     */
    private Set<ProxyInfo> proxies;

    public RegisterClientReq preCheck() {
        Assert.notBlank(token, "请输入token");

        final var registerClientReq = Cache.<RegisterClientReq>get(token);
        Assert.nonNull(registerClientReq, "token is not valid");

        Assert.notEmpty(proxies,"添加的代理信息为空");

        for (final ProxyInfo proxy : proxies) {
            proxy.preCheck();
        }

        return registerClientReq;
    }
}