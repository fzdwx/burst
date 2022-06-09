package server;

import core.Server;
import io.github.fzdwx.lambada.Seq;
import io.github.fzdwx.lambada.lang.StopWatch;
import io.netty.channel.nio.NioEventLoopGroup;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.TreeSet;

/**
 * @author <a href="mailto:likelovec@gmail.com">fzdwx</a>
 * @date 2022/5/22 12:14
 */
public class TestServer {

    // @Test
    // void name() {
    //     final var boss = new NioEventLoopGroup();
    //     final var worker = new NioEventLoopGroup();
    //     final var server = new Server()
    //             .withGroup(boss, worker)
    //             .onSuccess(s -> {
    //                 System.out.println(s.port());
    //             })
    //             .listen(0)
    //             .dispose();
    //
    // }

    @Test
    void test_seq() {
        for (int j = 0; j < 3; j++) {
            final var objects = new ArrayList<Integer>(10000000);
            for (int i = 0; i < 10000000; i++) {
                objects.add(i);
            }

            final var treeSet = new TreeSet<Long>();
            for (int i = 0; i < 100; i++) {
                final var task = StopWatch.get("test_seq:" + i);
                final var integers = Seq.of(objects).toList();
                treeSet.add(task.stop());
            }

            System.out.println("test seq:");
            System.out.println(treeSet.first());
            System.out.println(treeSet.last());
        }
    }

    @Test
    void test_stream() {
        for (int j = 0; j < 3; j++) {
            final var objects = new ArrayList<Integer>(10000000);
            for (int i = 0; i < 10000000; i++) {
                objects.add(i);
            }

            final var treeSet = new TreeSet<Long>();
            for (int i = 0; i < 100; i++) {
                final var task = StopWatch.get("test_seq:" + i);
                final var integers = objects.stream().toList();
                treeSet.add(task.stop());
            }
            System.out.println("test stream:");
            System.out.println(treeSet.first());
            System.out.println(treeSet.last());
        }
    }
}