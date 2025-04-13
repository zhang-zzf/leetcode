import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class NumberOfCommonFactorsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.commonFactors(12, 6)).isEqualTo(4);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int commonFactors(int a, int b) {
            int ans = 0;
            int max = Math.max(a, b);
            for (int i = 1; i <= max; i++) {
                if (a % i == 0 && b % i == 0) {
                    ans += 1;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}