package burst.server.logic.temp;

import io.github.fzdwx.lambada.Collections;

import java.util.Map;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 21:03
 */
public class Cache {

    private static final Map<String, String> cache = Collections.cMap();


    public static void set(String key, String value) {
        cache.put(key, value);
    }

    public static String get(String key) {
        return cache.get(key);
    }

    public static void remove(String key) {
        cache.remove(key);
    }
}