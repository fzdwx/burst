package burst.server.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/17 11:51
 */
@RestController
@RequestMapping("/test")
public class TestController {

    @GetMapping("/hello/{name}")
    public String hello(@PathVariable final String name) {
        return "hello world";
    }
}