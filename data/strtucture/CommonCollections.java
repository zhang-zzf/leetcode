import org.junit.jupiter.api.Test;

import java.util.*;

import static org.assertj.core.api.BDDAssertions.then;

class CommonCollections {

    public static final String A = "A";

    /**
     * List
     * can add null for ArrayList
     */
    @Test
    void givenList_when_thenSuccess() {
        List<String> list = new ArrayList<>();
        list.add(A);
        list.add(0, null);
        while (!list.isEmpty()) {
            then(list.remove(0)).isIn(null, A);
            then(list.remove(list.size() - 1)).isIn(null, A);
        }
    }

    /**
     * Stack
     * can add null for LinkedList
     */
    @Test
    void givenStack_when_thenSuccess() {
        Deque<String> stack = new LinkedList<>();
        // push
        stack.push(A);
        stack.push(null);
        // pop
        while (!stack.isEmpty()) {
            String last = stack.pop();
            then(last).isIn(null, A);
        }
    }

    /**
     * Queue
     * FIFO
     * can add null to the Queue
     */
    @Test
    void givenQueue_when_thenSuccess() {
        Queue<String> queue = new LinkedList<>();
        // enqueue
        queue.offer(A);
        queue.offer(null);
        // dequeue
        while (!queue.isEmpty()) {
            String pollHead = queue.poll();
            then(pollHead).isIn(null, A);
        }
        // peek head
        String peekHead = queue.peek();
        then(peekHead).isNull();
    }

    /**
     * Queue
     * FIFO
     * can add null to the queue
     */
    @Test
    void givenDeQueue_when_thenSuccess() {
        Deque<String> deque = new LinkedList<>();
        // tail
        deque.offerLast(A);
        deque.offerLast(null);
        while (!deque.isEmpty()) {
            String last = deque.pollLast();
            then(last).isIn(A, null);
        }
        // head
        deque.offerFirst(A);
        String head = deque.pollFirst();
        then(head).isEqualTo(A);
        // deque use as queue
        // poll from head, add to tail
        deque.offer(A);
        head = deque.poll();
        then(head).isEqualTo(A);
    }

    /**
     * PriorityQueue
     */
    @Test
    void givenPriorityQueue_when_thenSuccess() {
        PriorityQueue<int[]> pq = new PriorityQueue<>(Comparator.comparing(a -> a[0]));
        pq.offer(new int[]{1, 0});
        pq.offer(new int[]{2, 0});
        while (!pq.isEmpty()) {
            int[] peek = pq.peek();
            int[] poll = pq.poll();
        }
    }

}
