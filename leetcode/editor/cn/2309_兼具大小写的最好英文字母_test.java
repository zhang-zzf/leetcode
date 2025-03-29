import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class GreatestEnglishLetterInUpperAndLowerCaseTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.greatestLetter("lEeTcOdE")).isEqualTo("E");
        then(solution.greatestLetter("arRAzFif")).isEqualTo("R");
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String greatestLetter(String s) {
            String ans = "";
            int[] cnt = new int[26];
            for (char c : s.toCharArray()) {
                if (Character.isLowerCase(c)) {
                    cnt[c - 'a'] |= 0B01;
                }
                else {
                    cnt[c - 'A'] |= 0B10;
                }
            }
            for (int i = 25; i >= 0; i--) {
                if (cnt[i] == 0B11) {
                    ans = String.valueOf((char) ('A' + i));
                    break;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}