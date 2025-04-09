import static org.assertj.core.api.BDDAssertions.then;

import java.util.Comparator;
import java.util.PriorityQueue;
import org.junit.jupiter.api.Test;


class MinimumAmountOfTimeToFillCupsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.fillCups(new int[]{1, 4, 2})).isEqualTo(4);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int fillCups(int[] amount) {
            PriorityQueue<Integer> pq = new PriorityQueue<>(Comparator.reverseOrder());
            for (int num : amount) {
                if (num > 0) {
                    pq.offer(num);
                }
            }
            int ans = 0;
            while (!pq.isEmpty()) {
                Integer a = pq.poll();
                if (!pq.isEmpty()) {
                    Integer b = pq.poll();
                    if (b > 1) {
                        pq.offer(b - 1);
                    }
                    if (a > 1) {
                        pq.offer(a - 1);
                    }
                    ans += 1;
                }
                else {
                    ans += a;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}