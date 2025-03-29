import org.junit.jupiter.api.Test;


class CheckIfNumberHasEqualDigitCountAndDigitValueTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean digitCount(String num) {
            boolean ans = true;
            int[] cnt = new int[10];
            for (int i = 0; i < num.length(); i++) {
                cnt[num.charAt(i) - '0'] += 1;
            }
            for (int i = 0; i < num.length(); i++) {
                if (num.charAt(i) - '0' != cnt[i]) {
                    ans = false;
                    break;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}