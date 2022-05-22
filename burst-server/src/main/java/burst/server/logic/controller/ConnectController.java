package burst.server.logic.controller;

import burst.protocol.BurstFactory;
import burst.server.inf.redis.Redis;
import burst.server.logic.domain.model.request.RegisterInfo;
import burst.server.logic.trans.Transform;
import http.HttpServerRequest;
import io.github.fzdwx.lambada.Console;
import io.github.fzdwx.lambada.Exceptions;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdw1x</a>
 * @date 2022/5/21 17:09
 */
@RestController
@Slf4j
public class ConnectController {

    @GetMapping("connect")
    public void connect(@RequestParam String token, HttpServerRequest request) {
        final var registerInfo = RegisterInfo.from(Redis.get(token));
        if (registerInfo == null) {
            throw Exceptions.newIllegalArgument("token is invalid");
        }

        request.upgradeToWebSocket(ws -> {
            ws.mountOpen(h -> {
                final var portMap = Transform.init(registerInfo, ws, token);
                if (portMap == null) {
                    ws.sendBinary(BurstFactory.error("portMap is null,maybe server did not have available Port"));
                    return;
                }
                ws.sendBinary(BurstFactory.successForPort(portMap));
                final var reader = Console.defaultLineReader();
                while (true) {
                    ws.send(reader.readLine());
                }
            });

            ws.mountBinary(b -> {
                System.out.println(b);
            });
        });

    }
}