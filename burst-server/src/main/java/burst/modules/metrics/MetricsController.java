package burst.modules.metrics;

import burst.inf.metrics.MetricsRecorder;
import core.http.response.HttpResponse;
import core.http.response.JsonObjectHttpResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * metrics controller.
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/6/24 21:48
 */
@RestController
@RequestMapping("metrics")
@RequiredArgsConstructor
public class MetricsController {

    private final MetricsRecorder metricsRecorder;

    @GetMapping("{token}")
    public JsonObjectHttpResponse get(@PathVariable final String token) {
        return HttpResponse.json(metricsRecorder.get(token));
    }
}