package burst.domain.model.request;

import burst.domain.ProxyInfo;
import burst.temp.Cache;
import io.github.fzdwx.lambada.Assert;
import lombok.Data;

import java.util.Set;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/13 22:44
 */
@Data
public class RemoveProxyInfoReq {

    /**
     * 令牌
     */
    private String token;

    /**
     * 想要删除(关闭)的代理信息
     */
    private Set<ProxyInfo> proxies;

    public RegisterClientReq preCheck() {
        Assert.notBlank(token, "请输入token");

        final var registerClientReq = Cache.<RegisterClientReq>get(token);
        Assert.nonNull(registerClientReq, "token is not valid");

        Assert.notEmpty(proxies,"需要删除的代理信息为空");

        for (final ProxyInfo proxy : proxies) {
            proxy.preCheckForRemove();
        }

        return registerClientReq;
    }

}