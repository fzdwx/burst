package burst.modules.user.domain.model.request;

import burst.modules.user.domain.po.ProxyInfo;
import io.github.fzdwx.lambada.Collections;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Lang;
import io.github.fzdwx.lambada.Seq;
import lombok.Data;

import java.util.Collection;
import java.util.Set;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 17:01
 */
@Data
public class RegisterClientReq {

    public static final RegisterClientReq DEFAULT = new RegisterClientReq() {

        {
            final var element = new ProxyInfo();
            element.port = 63342;
            setProxies(Collections.set(element));
        }
    };

    /**
     * 需要被代理的内网地址信息
     */
    private Set<ProxyInfo> proxies;

    public void preCheck() {
        if (Lang.isEmpty(proxies)) {
            throw Exceptions.newIllegalArgument("port is empty");
        }

        for (final ProxyInfo proxy : proxies) {
            proxy.preCheck();
        }
    }

    /**
     * add all
     *
     * @apiNote 返回实际上添加成功了的
     */
    public Collection<ProxyInfo> addAll(final Set<ProxyInfo> proxies) {
        return Seq.of(proxies).filter(this.proxies::add).toList();
    }
}