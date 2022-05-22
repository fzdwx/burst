// package burst.server.inf.redis;
//
// import cn.hutool.extra.spring.SpringUtil;
// import io.github.fzdwx.lambada.anno.NonNull;
// import org.redisson.api.RedissonClient;
// import org.springframework.beans.factory.InitializingBean;
// import org.springframework.data.redis.core.HashOperations;
// import org.springframework.data.redis.core.ListOperations;
// import org.springframework.data.redis.core.SetOperations;
// import org.springframework.data.redis.core.StringRedisTemplate;
// import org.springframework.data.redis.core.ValueOperations;
// import org.springframework.stereotype.Component;
//
// import java.time.Duration;
// import java.util.Arrays;
// import java.util.Collection;
// import java.util.List;
// import java.util.Map;
//
// /**
//  * Redis.
//  *
//  * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
//  * @date 2021/12/30 13:58
//  */
// // @Component
// public final class Redis implements InitializingBean {
//
//     private static StringRedisTemplate REDIS_TEMPLATE;
//     private static ValueOperations<String, String> STRING;
//     private static HashOperations<String, String, String> HASH;
//     private static SetOperations<String, String> SET;
//     private static ListOperations<String, String> LIST;
//
//     private static RedissonClient redissonClient;
//
//
//     private Redis() {
//     }
//
//     /*
//                                              =======================================================
//                                              ==============                           ==============
//                                              ==============       redis keys          ==============
//                                              ==============                           ==============
//                                              =======================================================
//      */
//
//     /**
//      * 确定给定的key存在
//      */
//     public static Boolean exists(@NonNull final String key) {
//         return REDIS_TEMPLATE.hasKey(key);
//     }
//
//     /**
//      * 删除给定的key 。
//      *
//      * @return {@link Boolean } 如果删除了key，则为 true。
//      */
//     public static Boolean del(@NonNull final String key) {
//         return REDIS_TEMPLATE.delete(key);
//     }
//
//     /**
//      * 删除给定的keys
//      *
//      * @return {@link Long } 删除的键数。 在管道/事务中使用时为 null。
//      */
//     public static Long del(@NonNull final String... keys) {
//         return REDIS_TEMPLATE.delete(Arrays.asList(keys));
//     }
//
//     /**
//      * 删除给定的keys
//      *
//      * @return {@link Long } 删除的键数。 在管道/事务中使用时为 null。
//      */
//     public static Long del(@NonNull final Collection<String> key) {
//         return REDIS_TEMPLATE.delete(key);
//     }
//
//     /**
//      * 设置给定key生存时间。
//      *
//      * @return {@link Boolean }在管道/事务中使用时为 null。
//      * @throws IllegalArgumentException if the timeout is {@literal null}.
//      */
//     public static Boolean expire(@NonNull final String key, @NonNull final Duration timeout) {
//         return REDIS_TEMPLATE.expire(key, timeout);
//     }
//
//     /**
//      * ttl.
//      *
//      * @return {@link Long } 在管道/事务中使用时为 null。
//      */
//     public static Long ttl(@NonNull final String key) {
//         return REDIS_TEMPLATE.getExpire(key);
//     }
//
//     /*
//                                              =======================================================
//                                              ==============                             ============
//                                              ==============       redis string          ============
//                                              ==============                             ============
//                                              =======================================================
//      */
//
//     public static Long incr(final String key) {
//         return STRING.increment(key);
//     }
//
//     /**
//      * 为key设置value.
//      */
//     public static void set(@NonNull final String key, @NonNull final String val) {
//         STRING.set(key, val);
//     }
//
//     /**
//      * 设置key的value和过期timeout.
//      *
//      * @throws IllegalArgumentException 如果{@code key}, {@code value} or {@code timeout} 不存在。
//      */
//     public static void set(@NonNull final String key, @NonNull final String val, @NonNull final Duration timeout) {
//         STRING.set(key, val, timeout);
//     }
//
//     /**
//      * 如果key不存在，则设置key以保存字符串value 。
//      */
//     public static Boolean setNx(@NonNull final String key, @NonNull final String val) {
//         return STRING.setIfAbsent(key, val);
//     }
//
//     /**
//      * 如果key存在，则设置key以保存字符串value 。
//      */
//     public static Boolean setEx(@NonNull final String key, @NonNull final String val) {
//         return STRING.setIfPresent(key, val);
//     }
//
//     /**
//      * 如果存在key ，则设置key以保存字符串value和过期timeout 。
//      */
//     public static Boolean setEx(@NonNull final String key, @NonNull final String val, @NonNull final Duration timeout) {
//         return STRING.setIfPresent(key, val, timeout);
//     }
//
//     /**
//      * 如果key不存在，则设置key以保存字符串value和过期timeout 。
//      */
//     public static Boolean setNx(@NonNull final String key, @NonNull final String val, @NonNull final Duration timeout) {
//         return STRING.setIfAbsent(key, val, timeout);
//     }
//
//     /**
//      * 获取key的值。
//      *
//      * @return {@link String } 在管道/事务中使用时为 null。
//      */
//     public static String get(@NonNull final String key) {
//         return STRING.get(key);
//     }
//
//     /**
//      * 为key设置value,并返回旧值.
//      *
//      * @return 在管道/事务中使用时为 null。
//      */
//     public static String getSet(@NonNull final String key, @NonNull final String val) {
//         return STRING.getAndSet(key, val);
//     }
//
//     /**
//      * 获取存储在key中的值的长度。
//      *
//      * @return {@link Long } 在管道/事务中使用时为 null。
//      */
//     public static Long strLen(@NonNull final String key) {
//         return STRING.size(key);
//     }
//
//     /*
//                                              =======================================================
//                                              ==============                           ==============
//                                              ==============       redis  set          ==============
//                                              ==============                           ==============
//                                              =======================================================
//      */
//
//     /**
//      * 在key处添加给定的values 。
//      *
//      * @param key    key
//      * @param values values
//      * @return {@link Long } 成功添加的个数
//      */
//     public static Long sAdd(final String key, final String... values) {
//         return SET.add(key, values);
//     }
//
//     /**
//      * 从 set at key中删除给定values并返回已删除元素的数量。
//      *
//      * @param key    key
//      * @param values values
//      * @return {@link Long } 成功移除的个数
//      */
//     public static Long sRem(final String key, final Object... values) {
//         return SET.remove(key, values);
//     }
//
//     /*
//                                              =======================================================
//                                              ==============                           ==============
//                                              ==============       redis hash          ==============
//                                              ==============                           ==============
//                                              =======================================================
//      */
//
//     /**
//      * 设置哈希hashKey的value。
//      */
//     public static void hSet(@NonNull final String key, @NonNull final String field, final String val) {
//         HASH.put(key, field, val);
//     }
//
//     public static void hIncr(@NonNull final String key, @NonNull final String field, final long amount) {
//         HASH.increment(key, field, amount);
//     }
//
//     public static void hIncr(@NonNull final String key, @NonNull final String field) {
//         HASH.increment(key, field, 1);
//     }
//
//     /**
//      * 仅当hashKey不存在时才设置 hash hashKey的value。
//      */
//     public static void hSetNX(@NonNull final String key, @NonNull final String field, final String val) {
//         HASH.putIfAbsent(key, field, val);
//     }
//
//     /**
//      * 使用map提供的数据将多个哈希字段设置为多个值。。
//      */
//     public static void hmSet(@NonNull final String key, @NonNull final Map<String, String> map) {
//         HASH.putAll(key, map);
//     }
//
//     /**
//      * 删除给定的哈希 filed
//      *
//      * @param key    关键
//      * @param fields 字段
//      */
//     public static void hDel(@NonNull final String key, @NonNull final Object... fields) {
//         HASH.delete(key, fields);
//     }
//
//     /**
//      * 从key处的 hash 获取给定field的值。
//      *
//      * @return {@link String } 当 key 或 field 不存在或在管道/事务中使用时为 null。
//      */
//     public static String hGet(@NonNull final String key, @NonNull final String field) {
//         return HASH.get(key, field);
//     }
//
//     /**
//      * 从key处的 hash 获取给定fields的值。
//      *
//      * @return {@link List<String> } 在管道/事务中使用时为 null。
//      */
//     public static List<String> hmGet(@NonNull final String key, @NonNull final Collection<String> fields) {
//         return HASH.multiGet(key, fields);
//     }
//
//     /**
//      * 获取hash结构的所有数据
//      *
//      * @return {@link String } 在管道/事务中使用时为 null。
//      */
//     public static Map<String, String> hGetAll(@NonNull final String key) {
//         return HASH.entries(key);
//     }
//
//     /*
//                                              =======================================================
//                                              ==============                           ==============
//                                              ==============       redis hash          ==============
//                                              ==============                           ==============
//                                              =======================================================
//      */
//
//     public static List<String> lrange(String key, int start, int end) {
//         return LIST.range(key, start, end);
//     }
//
//     public static List<String> lrange(String key) {
//         return LIST.range(key, 0, -1);
//     }
//
//     public static void lpush(final String key, final String value) {
//         LIST.leftPush(key, value);
//     }
//
//         /*
//                                              =======================================================
//                                              ==============                           ==============
//                                              ==============       redis pub        ==============
//                                              ==============                           ==============
//                                              =======================================================
//         */
//
//     /**
//      * 将给定事件发布到给定主题
//      *
//      * @param topic 主题
//      * @param event 事件
//      */
//     public static void pub(@NonNull final String topic, @NonNull final Object event) {
//         REDIS_TEMPLATE.convertAndSend(topic, event);
//     }
//
//     @Override
//     public void afterPropertiesSet() throws Exception {
//         REDIS_TEMPLATE = SpringUtil.getBean("stringRedisTemplate");
//         redissonClient = SpringUtil.getBean(RedissonClient.class);
//         STRING = REDIS_TEMPLATE.opsForValue();
//         HASH = REDIS_TEMPLATE.opsForHash();
//         SET = REDIS_TEMPLATE.opsForSet();
//         LIST = REDIS_TEMPLATE.opsForList();
//     }
// }