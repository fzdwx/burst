package burst.modules.user.domain.model.request;

import burst.modules.user.domain.po.ProxyInfo;
import io.github.fzdwx.lambada.Collections;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Lang;
import io.github.fzdwx.lambada.Seq;
import lombok.Data;

import java.util.Collection;
import java.util.HashSet;
import java.util.List;
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
     * copy add all
     *
     * @apiNote 返回实际上添加成功了的
     */
    public Collection<ProxyInfo> copyAndAddAll(final Set<ProxyInfo> proxies) {
        if (Lang.isEmpty(proxies)) {
            return Collections.emptyList();
        }

        final HashSet<ProxyInfo> copy = new HashSet<>(proxies);
        return Seq.of(proxies).filter(copy::add).toList();
    }

    /**
     * copy remove all
     *
     * @apiNote 返回实际上删除成功了的
     */
    public List<ProxyInfo> copyRemoveAll(final Set<ProxyInfo> proxies) {
        if (Lang.isEmpty(proxies)) {
            return Collections.emptyList();
        }

        final HashSet<ProxyInfo> copy = new HashSet<>(proxies);
        return Seq.of(proxies).filter(copy::remove).toList();
    }

    /**
     * add all
     *
     * @param proxies proxy info
     * @return {@link RegisterClientReq }
     */
    public RegisterClientReq addAll(final Collection<ProxyInfo> proxies) {
        if (Lang.isNotEmpty(proxies)) {
            this.proxies.addAll(proxies);
        }

        return this;
    }

    /**
     * remove all
     *
     * @param proxies proxy info
     * @return {@link RegisterClientReq }
     */
    public RegisterClientReq removeAll(final List<ProxyInfo> proxies) {
        if (Lang.isNotEmpty(proxies)) {
            proxies.forEach(this.proxies::remove);
        }

        return this;
    }
}