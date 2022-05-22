// package burst.server.inf.redis.lock;
//
// import cn.hutool.extra.spring.SpringUtil;
// import io.github.fzdwx.lambada.Exceptions;
// import lombok.extern.slf4j.Slf4j;
// import org.redisson.api.RLock;
// import org.redisson.api.RedissonClient;
//
// import java.util.concurrent.TimeUnit;
// import java.util.function.Function;
//
// /**
//  * 分布式锁实现接口幂
//  *
//  * @author zhiyuan
//  * @since 2022/1/8
//  */
// @Slf4j
// public class ApiIdempotentSupport {
//
//     private static final int WAIT_LOCK = -1;
//     private static final int NOW = 0;
//     private static final String LOCK_KEY_PREFIX = "apiIdempotent_lock:";
//     private static final RedissonClient REDISSON_CLIENT;
//     private static final TimeUnit DEFAULT_TIME_UNIT = TimeUnit.SECONDS;
//
//     // init
//     static {
//         REDISSON_CLIENT = SpringUtil.getBean(RedissonClient.class);
//     }
//
//
//     public static <R> R idempotent(final String value, final int expireTime, final Function<Boolean, R> action) {
//         return idempotent(value, NOW, expireTime, action);
//     }
//
//     /**
//      * 加分布式锁(不解锁)
//      *
//      * @param value          锁的名字
//      * @param acquireTimeout 获取锁的等待时间
//      * @param expireTime     持有锁的时间
//      * @param action         你要做的事情，入参为是否加锁成功
//      * @return {@link R} 你返回的结果
//      */
//     public static <R> R idempotent(final String value, final int acquireTimeout, final int expireTime, final Function<Boolean, R> action) {
//         // init lock
//         final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(value));
//
//         // lock
//         try {
//             boolean flag = lock.tryLock(acquireTimeout, expireTime, DEFAULT_TIME_UNIT);
//             // do process
//             return action.apply(flag);
//         } catch (final InterruptedException ignore) {
//             lock.unlock();
//             throw Exceptions.newIllegalState("请稍后再试");
//         }
//     }
//
//
// }