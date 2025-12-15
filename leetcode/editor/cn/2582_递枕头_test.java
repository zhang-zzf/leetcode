import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class PassThePillowTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.passThePillow(4, 5)).isEqualTo(2);
        then(solution.passThePillow(3, 2)).isEqualTo(3);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int passThePillow(int n, int time) {
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