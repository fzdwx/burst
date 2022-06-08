package burst.modules.user.controller;

import burst.modules.user.domain.model.request.RegisterInfo;
import burst.temp.Cache;
import cn.hutool.core.util.IdUtil;
import io.github.fzdwx.lambada.Collections;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;


/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 16:54
 */
@RestController
public class RegisterController {

    /**
     * 注册客户端
     */
    @PostMapping("register")
    public ResponseEntity<?> register(@RequestBody RegisterInfo info) {
        info.preCheck();

        final String token = IdUtil.fastSimpleUUID();
        Cache.set(token, info.encode());

        return ResponseEntity.ok().body(Collections.map(
                "token", token
        ));
    }
}