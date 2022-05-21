package burst.server.inf.redis.lock;

import cn.hutool.extra.spring.SpringUtil;
import io.github.fzdwx.lambada.Exceptions;
import io.github.fzdwx.lambada.fun.Callable;
import lombok.extern.slf4j.Slf4j;
import org.redisson.api.RLock;
import org.redisson.api.RedissonClient;

import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;
import java.util.function.Function;

/**
 * 分布式锁的方法实现
 * <pre>
 *     public String get() {
 *          final String key ...;
 *          final String jsonStr = this.redisService.String.get(key);
 *          if (StringUtils.isNull(jsonStr)) {
 *              return this.getData(key);
 *          }
 *          return jsonStr;
 *     }
 *
 *     private String getData(key) {
 *          return DistributedLockSupport.lock(key, 10, (bool) -> {
 *              if (bool) {
 *                  return "12311123d";
 *              } else return this.get(); // 没有获取到锁，从头开始
 *          });
 *     }
 * </pre>
 *
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2021/12/2 12:08
 */
@Slf4j
public class DistributedLockSupport {

    private static final int WAIT_LOCK = -1;
    private static final int NOW = 0;
    private static final String LOCK_KEY_PREFIX = "distributed_lock:";
    private static final RedissonClient REDISSON_CLIENT;
    private static final TimeUnit DEFAULT_TIME_UNIT = TimeUnit.SECONDS;

    // init
    static {
        REDISSON_CLIENT = SpringUtil.getBean(RedissonClient.class);
    }

    /**
     * 加分布式锁
     *
     * @param lockName 锁的名字
     * @param action   你要做的事情，入参为是否加锁成功
     * @return {@link R} 你返回的结果
     * @apiNote 立即执行tryLock, 如果获取到锁则当执行完时, 自动释放锁;否则抛出异常
     */
    public static <R> R tryLock(final String lockName, final Function<Boolean, R> action) {
        return tryLock(lockName, NOW, WAIT_LOCK, action);
    }

    /**
     * 加分布式锁
     *
     * @param lockName   锁的名字
     * @param expireTime 持有锁的时间
     * @param action     你要做的事情，入参为是否加锁成功
     * @return {@link R} 你返回的结果
     * @apiNote 立即执行tryLock, 如果获取到锁则expireTime后自动释放锁, 不论是否调用unlock释放锁;否则抛出异常
     * 单位固定为秒
     */
    public static <R> R tryLock(final String lockName, final int expireTime, final Function<Boolean, R> action) {
        return tryLock(lockName, NOW, expireTime, action);
    }

    /**
     * 加分布式锁
     *
     * @param lockName       锁的名字
     * @param acquireTimeout 获取锁的等待时间
     * @param expireTime     持有锁的时间
     * @param action         你要做的事情，入参为是否加锁成功
     * @return {@link R} 你返回的结果
     * @apiNote 立即执行tryLock, 最多阻塞等待acquireTimeout秒, 如果获取到锁则等待expireTime秒后自动释放锁, 不论是否调用unlock释放锁;否则抛出异常
     * 单位固定为秒
     */
    public static <R> R tryLock(final String lockName, final int acquireTimeout, final int expireTime, final Function<Boolean, R> action) {
        // init lock
        final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(lockName));

        // lock
        boolean flag = false;
        try {
            flag = lock.tryLock(acquireTimeout, expireTime, DEFAULT_TIME_UNIT);
            // do process
            return action.apply(flag);
        } catch (final InterruptedException ignore) {
            throw Exceptions.newIllegalState("请稍后再试");
        } finally {
            try {
                if (flag) {
                    lock.unlock();
                }
            } catch (final Exception e) {
                log.error("RLock unlock error:{}", e.getMessage(), e);
            }
        }
    }

    /**
     * 加分布式锁
     *
     * @param lockName 锁的名字
     * @param action   你要做的事情，入参为是否加锁成功
     * @apiNote 立即执行tryLock, 如果获取到锁则当执行完时, 自动释放锁;否则抛出异常
     */
    public static void tryLock(final String lockName, final Consumer<Boolean> action) {
        tryLock(lockName, NOW, WAIT_LOCK, action);
    }

    /**
     * 加分布式锁
     *
     * @param lockName   锁的名字
     * @param expireTime 持有锁的时间
     * @param action     你要做的事情，入参为是否加锁成功
     * @apiNote 立即执行tryLock, 如果获取到锁则expireTime后自动释放锁, 不论是否调用unlock释放锁;否则抛出异常.
     * 单位固定为秒
     */
    public static void tryLock(final String lockName, final int expireTime, final Consumer<Boolean> action) {
        tryLock(lockName, NOW, expireTime, action);
    }

    /**
     * 加分布式锁
     *
     * @param lockName       锁的名字
     * @param acquireTimeout 获取锁的等待时间
     * @param expireTime     持有锁的时间
     * @param action         你要做的事情，入参为是否加锁成功
     * @apiNote 立即执行tryLock, 最多阻塞等待acquireTimeout秒, 如果获取到锁则等待expireTime秒后自动释放锁, 不论是否调用unlock释放锁;否则抛出异常.
     * 单位固定为秒
     */
    public static void tryLock(final String lockName, final int acquireTimeout, final int expireTime, final Consumer<Boolean> action) {
        // init lock
        final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(lockName));

        // lock
        boolean locked = false;
        try {
            locked = lock.tryLock(acquireTimeout, expireTime, DEFAULT_TIME_UNIT);
            // do process
            action.accept(locked);
        } catch (final InterruptedException e) {
            throw Exceptions.newIllegalState("请稍后再试");
        } finally {
            try {
                if (locked) {
                    lock.unlock();
                }
            } catch (final Exception e) {
                log.error("RLock unlock error:{}", e.getMessage(), e);
            }
        }
    }

    /**
     * 加分布式锁 保证执行
     *
     * @param lockName 锁的名字
     * @param action   你要做的事情，入参为是否加锁成功
     * @apiNote 如果锁不可用，则当前线程将被禁用以用于线程调度目的并处于休眠状态，直到获得锁为止。
     */
    public static void lock(final String lockName, final Runnable action) {
        // init lock
        final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(lockName));

        // lock
        try {
            lock.lock();
            // do process
            action.run();
        } finally {
            try {
                lock.unlock();
            } catch (final Exception e) {
                log.error("RLock unlock error:{}", e.getMessage(), e);
            }
        }
    }

    /**
     * 加分布式锁 保证执行
     *
     * @param lockName   锁的名字
     * @param expireTime 持有锁的时间
     * @param action     你要做的事情，入参为是否加锁成功
     * @apiNote 单位固定为秒, 使用定义的leaseTime获取锁。如有必要，等待锁定可用。锁定将在定义leaseTime时间间隔后自动释放。
     * 如果leaseTime 为-1，则保持锁定直到显式解锁。
     */
    public static void lock(final String lockName, final int expireTime, final Runnable action) {
        // init lock
        final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(lockName));

        // lock
        try {
            lock.lock(expireTime, DEFAULT_TIME_UNIT);
            // do process
            action.run();
        } finally {
            try {
                lock.unlock();
            } catch (final Exception e) {
                log.error("RLock unlock error:{}", e.getMessage(), e);
            }
        }
    }

    /**
     * 加分布式锁 保证执行
     *
     * @param lockName   锁的名字
     * @param expireTime 持有锁的时间
     * @param action     你要做的事情，入参为是否加锁成功
     * @return R
     * @apiNote 使用定义的expireTime获取锁。如有必要，等待锁定可用。锁定将在定义expireTime时间间隔后自动释放.
     * 单位固定为秒
     */
    public static <R> R lock(final String lockName, final int expireTime, final Callable<R> action) throws Exception {
        // init lock
        final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(lockName));

        try {
            lock.lock(expireTime, DEFAULT_TIME_UNIT);
            return action.call();
        } finally {
            try {
                lock.unlock();
            } catch (final Exception e) {
                log.error("RLock unlock error:{}", e.getMessage(), e);
            }
        }
    }

    /**
     * 加分布式锁 保证执行
     *
     * @param lockName 锁的名字
     * @param action   你要做的事情，入参为是否加锁成功
     * @return R
     * @apiNote 如果锁不可用，则当前线程将被禁用以用于线程调度目的并处于休眠状态，直到获得锁为止。
     */
    public static <R> R lock(final String lockName, final Callable<R> action) {
        // init lock
        final RLock lock = REDISSON_CLIENT.getLock(LOCK_KEY_PREFIX.concat(lockName));

        lock.lock();
        try {
            return action.call();
        } finally {
            try {
                lock.unlock();
            } catch (final Exception e) {
                log.error("RLock unlock error:{}", e.getMessage(), e);
            }
        }
    }
}