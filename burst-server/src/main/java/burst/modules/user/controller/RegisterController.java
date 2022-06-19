package burst.modules.user.controller;

import burst.domain.model.request.AddProxyInfoReq;
import burst.domain.model.request.RegisterClientReq;
import burst.domain.model.request.RemoveProxyInfoReq;
import burst.modules.user.service.RegisterService;
import burst.temp.Cache;
import core.http.response.HttpResponse;
import core.http.response.JsonMapHttpResponse;
import core.http.response.JsonObjectHttpResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 16:54
 */
@RestController
@RequestMapping("register")
@RequiredArgsConstructor
public class RegisterController {

    private final RegisterService registerService;

    /**
     * 注册客户端
     */
    @PostMapping
    public JsonMapHttpResponse register(@RequestBody RegisterClientReq req) {
        return HttpResponse.json("token", registerService.register(req));
    }

    /**
     * 添加代理信息
     */
    @PostMapping("addProxyInfo")
    public HttpResponse<?> addProxyInfo(@RequestBody AddProxyInfoReq req) {
        return HttpResponse.ok().body(() -> registerService.addProxyInfo(req));
    }

    /**
     * 删除代理信息（动态删除客户端需要代理的端口）
     *
     * @param req 要求事情
     * @return {@link JsonObjectHttpResponse }
     */
    @PostMapping("removeProxyInfo")
    public JsonObjectHttpResponse removeProxyInfo(@RequestBody RemoveProxyInfoReq req) {
        return HttpResponse.json(registerService.removeProxyInfo(req));
    }

    /**
     * 获取对应token的代理信息
     *
     * @param token 令牌
     * @return {@link JsonObjectHttpResponse }
     */
    @GetMapping("getProxyInfo")
    public JsonObjectHttpResponse getProxyInfo(@RequestParam String token) {
        return HttpResponse.json(Cache.<RegisterClientReq>get(token));
    }
}