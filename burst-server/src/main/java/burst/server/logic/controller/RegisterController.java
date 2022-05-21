package burst.server.logic.controller;

import burst.server.inf.redis.Redis;
import burst.server.logic.domain.model.request.RegisterReq;
import cn.hutool.core.util.IdUtil;
import http.HttpServerRequest;
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
        public ResponseEntity<?> register(@RequestBody RegisterReq req) {
        final String key = req.toKey();
        final String token = IdUtil.fastSimpleUUID();

        if (Redis.setNx(key, token)) {
            return ResponseEntity.ok().body(Collections.map(
                    "token", token
            ));
        }

        return ResponseEntity.badRequest().body("已经注册过了");
    }
}