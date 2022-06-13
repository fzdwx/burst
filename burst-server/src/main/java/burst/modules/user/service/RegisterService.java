package burst.modules.user.service;

import burst.modules.connect.controller.trans.Transform;
import burst.modules.user.domain.model.request.AddProxyInfoReq;
import burst.modules.user.domain.model.request.RegisterClientReq;
import burst.modules.user.domain.model.request.RemoveProxyInfoReq;
import burst.modules.user.domain.po.ProxyInfo;
import burst.temp.Cache;
import cn.hutool.core.util.IdUtil;
import io.github.fzdwx.lambada.Assert;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/12 13:46
 */
@Service
public class RegisterService {

    /**
     * 注册客户端(比如说提供客户端想要被代理的ip)
     *
     * @param req req
     * @return {@link String } Token
     */
    public String register(final RegisterClientReq req) {
        req.preCheck();

        final String token = IdUtil.fastSimpleUUID();
        Cache.set(token, req);

        return token;
    }

    /**
     * 添加代理信息（动态添加需要代理的端口）
     */
    public void addProxyInfo(final AddProxyInfoReq req) {
        final var registerClientReq = req.preCheck();

        final var proxies = registerClientReq.addAll(req.getProxies());
        Assert.notEmpty(proxies, "暂无需要代理的端口(您输入的代理信息可能已经存在！)");

        Transform.addProxyInfo(req.getToken(), req.getProxies());

        // 到时候可能不是内存缓存，所以需要更新
        Cache.set(req.getToken(), registerClientReq);
    }

    /**
     * 删除代理信息（动态删除客户端需要代理的端口）
     */
    public List<ProxyInfo> removeProxyInfo(final RemoveProxyInfoReq req) {
        final var registerClientReq = req.preCheck();

        final var proxies = registerClientReq.removeAll(req.getProxies());
        if (proxies.isEmpty()) {
            return null;
        }

        Transform.removeProxyInfo(req.getToken(), req.getProxies());

        // 到时候可能不是内存缓存，所以需要更新
        Cache.set(req.getToken(), registerClientReq);
        return proxies;
    }
}