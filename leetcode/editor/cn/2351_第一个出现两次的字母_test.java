import org.junit.jupiter.api.Test;


class FirstLetterToAppearTwiceTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public char repeatedCharacter(String s) {
            int[] cnt = new int[27];
            for (char c : s.toCharArray()) {
                cnt[c - 'a'] += 1;
                if (cnt[c - 'a'] >= 2) {
                    return c;
                }
            }
            return ' ';
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}