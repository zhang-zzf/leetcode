import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class NumberOfValidClockTimesTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.countTime("?5:00")).isEqualTo(2);
        then(solution.countTime("0?:0?")).isEqualTo(100);
        then(solution.countTime("??:??")).isEqualTo(1440);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countTime(String time) {
            int ans = 1;
            if (time.charAt(0) == '?') {
                char h2 = time.charAt(1);
                if (h2 == '?') {
                    ans = 24;
                }
                else {
                    ans *= h2 <= '3' ? 3 : 2;
                }
            }
            if (time.charAt(1) == '?') {
                char h1 = time.charAt(0);
                if (h1 == '?') {
                    ans = 24;
                }
                else if (h1 == '2') {
                    ans *= 4;
                }
                else {
                    ans *= 10;
                }
            }
            if (time.charAt(3) == '?') {
                ans *= 6;
            }
            if (time.charAt(4) == '?') {
                ans *= 10;
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}