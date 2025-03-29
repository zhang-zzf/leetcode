import org.junit.jupiter.api.Test;


class PercentageOfLetterInStringTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int percentageLetter(String s, char letter) {
            int ans = 0;
            for (char c : s.toCharArray()) {
                if (c == letter) {
                    ans += 1;
                }
            }
            return ans * 100 / s.length();
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}