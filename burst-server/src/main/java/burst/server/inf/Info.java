package burst.server.inf;

import com.alibaba.fastjson2.JSON;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 10:12
 */
public interface Info {

    default String encode() {
        return JSON.toJSONString(this);
    }
}