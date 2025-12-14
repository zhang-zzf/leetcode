import static org.assertj.core.api.BDDAssertions.then;

import java.util.PriorityQueue;
import org.junit.jupiter.api.Test;


class TakeGiftsFromTheRichestPileTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        long picked = solution.pickGifts(new int[]{25, 64, 9, 4, 100}, 4);
        then(picked).isEqualTo(29);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public long pickGifts(int[] gifts, int k) {
            return v2(gifts, k);
        }

        private long v2(int[] gifts, int k) {
            PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> b - a);
            for (int gift : gifts) {
                pq.add(gift);
            }
            while (k > 0 && !pq.isEmpty()) {
                int cur = pq.poll();
                pq.offer((int) Math.sqrt(cur));
                if (cur == 1) { // 最大值为1，没必要再取
                    break;
                }
                k--;
            }
            long sum = 0;
            while (!pq.isEmpty()) {
                sum += pq.poll();
            }
            return sum;
        }

        // sum int 越界
        private int failedCase1(int[] gifts, int k) {
            int sum = 0;
            PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> b - a);
            for (int gift : gifts) {
                sum += gift;
                pq.add(gift);
            }
            while (k > 0 && !pq.isEmpty()) {
                int cur = pq.poll();
                if (cur == 1) { // 最大值为1，没必要再取
                    break;
                }
                int left = (int) Math.sqrt(cur);
                sum -= cur - left;
                pq.offer(left);
                k--;
            }
            return sum;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}