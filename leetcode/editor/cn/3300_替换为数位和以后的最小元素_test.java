import org.junit.jupiter.api.Test;


class MinimumElementAfterReplacementWithDigitSumTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int minElement(int[] nums) {
            int ans = Integer.MAX_VALUE;
            for (int num : nums) {
                int sum = 0;
                while (num > 0) {
                    sum += num % 10;
                    num /= 10;
                }
                if (sum < ans) {
                    ans = sum;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}