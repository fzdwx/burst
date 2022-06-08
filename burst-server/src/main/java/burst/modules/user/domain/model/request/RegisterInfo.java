package burst.modules.user.domain.model.request;

import burst.inf.Info;
import com.alibaba.fastjson2.JSON;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.Lang;
import io.github.fzdwx.lambada.anno.Nullable;
import lombok.Data;

import java.util.List;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/21 17:01
 */
@Data
public class RegisterInfo implements Info {

    /**
     * 需要暴露的端口号
     */
    private List<Integer> ports;

    public void preCheck() {
        if (Lang.isEmpty(ports)) {
            throw Exceptions.newIllegalArgument("port is empty");
        }
    }

    @Nullable
    public static RegisterInfo from(String jsonStr) {
        if (Lang.isEmpty(jsonStr)) {
            return null;
        }

        return JSON.parseObject(jsonStr, RegisterInfo.class);
    }
}