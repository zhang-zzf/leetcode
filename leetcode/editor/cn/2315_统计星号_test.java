import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CountAsterisksTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.countAsterisks("l|*e*et|c**o|*de|")).isEqualTo(2);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countAsterisks(String s) {
            int ans = 0;
            for (int i = 0, cnt = 0; i < s.length(); i++) {
                char ch = s.charAt(i);
                if (ch == '|') {
                    cnt++;
                }
                if (ch == '*' && cnt % 2 == 0) {
                    ans++;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}