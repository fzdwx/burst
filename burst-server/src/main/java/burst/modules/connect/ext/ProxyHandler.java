package burst.modules.connect.ext;

import burst.domain.ProxyInfo;
import burst.domain.ServerUserConnectContainer;
import core.Server;
import io.github.fzdwx.lambada.anno.NonNull;
import io.github.fzdwx.lambada.anno.Nullable;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/18 21:31
 */
public interface ProxyHandler {

    /**
     * can handle proxy type
     *
     * @return {@link burst.domain.ProxyType }
     */
    @NonNull
    String supportType();

    @Nullable
    Server apply(String token, ServerUserConnectContainer container, final ProxyInfo proxyInfo);
}