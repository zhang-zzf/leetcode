import org.junit.jupiter.api.Test;


class MinimumCutsToDivideACircleTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int numberOfCuts(int n) {
            if (n == 1) {
                return 0;
            }
            return n % 2 == 0 ? n / 2 : n;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}