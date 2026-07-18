import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class PassThePillowTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.passThePillow(4, 5)).isEqualTo(2);
        then(solution.passThePillow(3, 2)).isEqualTo(3);
        then(solution.passThePillow(18, 38)).isEqualTo(5);
        then(solution.passThePillow(1000, 1000)).isEqualTo(999);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int passThePillow(int n, int time) {
            // return solution1(n, time);
            int idx = (time) % (n + n - 2);
            return idx < n ? idx + 1 : (n - (idx - (n - 1)));
        }

        private static int solution1(int n, int time) {
            int cur = 1, step = 1;
            for (int i = time; i > 0; i--) {
                if (cur == 1) {
                    step = 1;
                }
                if (cur == n) {
                    step = -1;
                }
                cur += step;
            }
            return cur;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}