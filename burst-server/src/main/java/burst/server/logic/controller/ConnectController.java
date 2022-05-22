package burst.server.logic.controller;

import burst.server.inf.redis.Redis;
import burst.server.logic.domain.model.request.RegisterInfo;
import core.Netty;
import http.HttpServerRequest;
import io.github.fzdwx.lambada.Exceptions;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdw1x</a>
 * @date 2022/5/21 17:09
 */
@RestController
public class ConnectController {

    @GetMapping("connect")
    public void connect(@RequestParam String token, HttpServerRequest request) {
        final var registerInfo = RegisterInfo.from(Redis.get(token));
        if (registerInfo == null) {
            throw Exceptions.newIllegalArgument("token is invalid");
        }

        request.upgradeToWebSocket(ws -> {
            ws.mountBinary(b -> {
                System.out.println(Netty.read(b));
            });
        });
    }
}