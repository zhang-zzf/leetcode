import org.junit.jupiter.api.Test;

import java.util.PriorityQueue;

import static org.assertj.core.api.BDDAssertions.then;


class MakeArrayZeroBySubtractingEqualAmountsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minimumOperations(new int[]{1, 5, 0, 3, 5})).isEqualTo(3);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int minimumOperations(int[] nums) {
            int ans = 0;
            int[] cnt = new int[101];
            for (int num : nums) {
                cnt[num] += 1;
            }
            for (int i = 1; i < cnt.length; i++) {
                if (cnt[i] > 0) {
                    ans += 1;
                }
            }
            return ans;
        }

        private static int wrong(int[] nums) {
            int ans = 0;
            PriorityQueue<Integer> pq = new PriorityQueue<>();
            for (int num : nums) {
                if (num > 0) {
                    pq.add(num);
                }
            }
            while (!pq.isEmpty()) {
                Integer cur = pq.poll();
                ans += 1;
                // 算法是错误的
                for (int i = 0; i < pq.size(); i++) {
                    Integer next = pq.poll();
                    if (next - cur > 0) {
                        pq.add(next - cur);
                    }
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}