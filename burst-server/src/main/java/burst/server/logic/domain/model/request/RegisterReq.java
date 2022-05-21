package burst.server.logic.domain.model.request;

import lombok.Data;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 17:01
 */
@Data
public class RegisterReq {

    private String ip;
    private String port;

    private static String PREFIX = "reg:";

    public String toKey() {
        return PREFIX + ip + port;
    }
}