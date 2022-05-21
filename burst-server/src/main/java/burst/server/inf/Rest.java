package burst.server.inf;

import cn.hutool.core.util.ArrayUtil;
import org.jetbrains.annotations.NotNull;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.util.LinkedHashMap;
import java.util.function.Supplier;

/**
 * rest response.
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/4/4 19:06
 */
public class Rest<OUT> extends ResponseEntity<Rest.Info<OUT>> {

    public static final String DATA = "data";
    public static final String CODE = "code";
    public static final String MESSAGE = "message";
    public static final String STACKTRACE = "stackTrace";

    public static final int SUCCESS = 0;
    public static final String SUCCESS_MESSAGE = "ok";
    public static final int FAILURE = 20001;
    public static final String FAILURE_MESSAGE = "failure";

    public static final int verify = 20010;

    public Rest(final Info<OUT> body, final HttpStatus status) {
        super(body, status);
    }

    public static <OUT> Rest<OUT> unauthorized(final String message) {
        return unauthorized(message, null);
    }

    public static <OUT> Rest<OUT> unauthorized(final String message, final StackTraceElement[] stackTrace) {
        return failure(message, HttpStatus.UNAUTHORIZED.value(), stackTrace);
    }

    public static <OUT> Rest<OUT> forbidden(final String message, final StackTraceElement[] stackTrace) {
        return failure(message, HttpStatus.FORBIDDEN.value(), stackTrace);
    }

    public static <OUT> Rest<OUT> verify(final String message, final StackTraceElement[] stackTrace) {
        return failure(message, verify, stackTrace);
    }

    public static <OUT> Rest<OUT> failure() {
        return failure(FAILURE_MESSAGE);
    }

    public static <OUT> Rest<OUT> failure(String message) {
        return failure(message, FAILURE);
    }

    public static <OUT> Rest<OUT> failure(final String message, final int code) {
        return failure(message, code, null);
    }

    public static <OUT> Rest<OUT> failure(String message, StackTraceElement[] stackTrace) {
        return failure(message, FAILURE, stackTrace);
    }

    public static <OUT> Rest<OUT> failure(final String message, int code, final StackTraceElement[] stackTrace) {
        return create(null, HttpStatus.INTERNAL_SERVER_ERROR, message, code, stackTrace);
    }

    @NotNull
    public static <OUT> Rest<OUT> success() {
        return create(null, HttpStatus.OK, SUCCESS_MESSAGE, SUCCESS, null);
    }

    @NotNull
    public static <OUT> Rest<OUT> success(OUT out) {
        return success(out, SUCCESS_MESSAGE);
    }

    @NotNull
    public static <OUT> Rest<OUT> success(OUT data, String message) {
        return create(data, HttpStatus.OK, message, SUCCESS, null);
    }

    @NotNull
    public static <OUT> Rest<OUT> success(Supplier<OUT> sup) {
        return success(sup.get());
    }

    @NotNull
    public static <OUT> Rest<OUT> success(Supplier<OUT> sup, String message) {
        return success(sup.get(), message);
    }

    @NotNull
    public static <OUT> Rest<OUT> of(final OUT data) {
        if (data instanceof Boolean) {
            if (data == Boolean.TRUE) {
                return success();
            } else return failure();
        }
        return success(data);
    }

    @NotNull
    public static <OUT> Rest<OUT> of(final Runnable action) {
        action.run();
        return success();
    }

    @NotNull
    public static <OUT> Rest<OUT> create(OUT data, HttpStatus status, String message, int code, StackTraceElement[] stackTrace) {
        final var outInfo = new Info<OUT>();
        if (message != null) {
            outInfo.put(CODE, code);
        }

        if (data != null) {
            outInfo.put(DATA, data);
        }

        if (message != null) {
            outInfo.put(MESSAGE, message);
        }

        if (stackTrace != null) {
            outInfo.put(STACKTRACE, ArrayUtil.sub(stackTrace, 0, 5));
        }

        return new Rest<>(outInfo, status);
    }

    public static class Info<OUT> extends LinkedHashMap<String, Object> {

    }
}