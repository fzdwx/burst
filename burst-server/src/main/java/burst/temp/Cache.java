package burst.temp;

import burst.modules.user.domain.model.request.RegisterClientReq;
import io.github.fzdwx.lambada.Collections;

import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 21:03
 */
public class Cache {

    private static final Map<String, Object> cache = Collections.cMap("b92a205269d94d38808c3979615245eb", RegisterClientReq.DEFAULT);


    public static void set(String key, Object value) {
        cache.put(key, value);
    }

    public static <T> T get(String key) {
        return (T) cache.get(key);
    }

    public static void remove(String key) {
        cache.remove(key);
    }
}