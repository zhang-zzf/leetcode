import org.junit.jupiter.api.Test;


class CountOperationsToObtainZeroTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countOperations(int num1, int num2) {
            int ans = 0;
            for (; num1 != 0 && num2 != 0; ans++) {
                if (num1 > num2) {
                    num1 = num1 - num2;
                }
                else {
                    num2 = num2 - num1;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}