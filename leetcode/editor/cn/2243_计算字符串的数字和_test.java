import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CalculateDigitSumOfAStringTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.digitSum("11111222223", 3)).isEqualTo("135");
        then(solution.digitSum("1234", 2)).isEqualTo("37");
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String digitSum(String s, int k) {
            if (s.length() <= k) {
                return s;
            }
            StringBuilder buf = new StringBuilder();
            for (int i = 0, j = 0, sum = 0; i < s.length(); i++) {
                sum += (s.charAt(i) - '0');
                // 注意最后一个分段 i == s.length() - 1
                if (++j == k || i == s.length() - 1) {
                    j = 0;
                    buf.append(sum);
                    sum = 0;
                }
            }
            return digitSum(buf.toString(), k);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}