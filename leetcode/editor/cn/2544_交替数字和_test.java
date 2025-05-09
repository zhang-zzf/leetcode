import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class AlternatingDigitSumTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.alternateDigitSum(5218)).isEqualTo(-4);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int alternateDigitSum(int n) {
            int ans = 0, sign = 1;
            while (n > 0) {
                ans += sign * (n % 10);
                n /= 10;
                sign *= -1;
            }
            return sign == -1 ? ans : -ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}