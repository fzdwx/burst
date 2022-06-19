package burst.modules.connect.controller;

import burst.modules.connect.trans.Transform;
import burst.domain.model.request.RegisterClientReq;
import burst.protocol.BurstMessage;
import burst.temp.Cache;
import com.google.protobuf.InvalidProtocolBufferException;
import core.http.ext.HttpServerRequest;
import core.http.ext.HttpServerResponse;
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
    public void connect(@RequestParam String token, HttpServerRequest request, HttpServerResponse response) {
        final var registerInfo = Cache.<RegisterClientReq>get(token);
        if (registerInfo == null) {
            response.end("token无效,请确认是否注册");
            return;
        }

        request.upgradeToWebSocket(ws -> {

            // step 1 [init] server export ports and send ports mapping to client.
            ws.mountOpen(h -> {
                Transform.init(registerInfo, ws, token);
            });

            ws.mountClose(h -> {
                Transform.destroy(token);
            });

            ws.mountBinary(b -> {
                BurstMessage burstMessage = null;
                try {
                    burstMessage = BurstMessage.parseFrom(b);
                } catch (InvalidProtocolBufferException e) {
                    log.error("parseFrom error", e);
                }

                if (burstMessage == null) {
                    return;
                }

                switch (burstMessage.getType()) {
                    // step 6 [forward to user]
                    case FORWARD_DATA -> Transform.toUser(burstMessage);
                    default -> log.error("unknown type {}", burstMessage.getType());
                }

            });

        });

    }
}