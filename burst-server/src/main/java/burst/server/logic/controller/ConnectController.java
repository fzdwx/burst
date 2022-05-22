package burst.server.logic.controller;

import burst.protocol.BurstFactory;
import burst.protocol.BurstMessage;
import burst.protocol.BurstType;
import burst.server.inf.redis.Redis;
import burst.server.logic.domain.model.request.RegisterInfo;
import burst.server.logic.trans.Transform;
import com.google.protobuf.InvalidProtocolBufferException;
import core.Netty;
import http.HttpServerRequest;
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

            // step 1 [init] server export ports and send ports mapping to client.
            ws.mountOpen(h -> {
                final var portMap = Transform.init(registerInfo, ws, token);
                if (portMap == null) {
                    ws.sendBinary(BurstFactory.error(BurstType.INIT,
                            "portMap is null,maybe server did not have available Port"));
                    return;
                }

                ws.sendBinary(BurstFactory.successForPort(portMap));
            });

            ws.mountBinary(b -> {
                BurstMessage burstMessage = null;
                try {
                    burstMessage = BurstMessage.parseFrom(Netty.readBytes(b));
                } catch (InvalidProtocolBufferException e) {
                    log.error("parseFrom error", e);
                }

                if (burstMessage == null) {
                    return;
                }

                switch (burstMessage.getType()) {
                    // step 6 [forward to user]
                    case FORWARD_DATA -> Transform.transform(burstMessage);
                    default -> log.error("unknown type {}", burstMessage.getType());
                }

            });

        });

    }
}