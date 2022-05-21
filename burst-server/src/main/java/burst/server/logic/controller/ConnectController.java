package burst.server.logic.controller;

import http.HttpServerRequest;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 17:09
 */
@RestController
public class ConnectController {

    @GetMapping("connect")
    public void connect(@RequestParam int token, HttpServerRequest request) {
        System.out.println("token = " + token);
        System.out.println(request);
    }

}