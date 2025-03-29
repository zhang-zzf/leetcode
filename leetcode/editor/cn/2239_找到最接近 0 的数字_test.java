import org.junit.jupiter.api.Test;


class FindClosestNumberToZeroTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int findClosestNumber(int[] nums) {
            int ans = Integer.MAX_VALUE;
            for (int num : nums) {
                if (Math.abs(num) < Math.abs(ans)) {
                    ans = num;
                }
                else if (Math.abs(num) == Math.abs(ans)) {
                    ans = Math.max(ans, num);
                }
            }
            return ans;
        }

        // failed case: [-1,-1]
        private static int v1(int[] nums) {
            int ans = Integer.MAX_VALUE;
            for (int num : nums) {
                if (num < 0) {
                    num *= -1;
                }
                if (num < ans) {
                    ans = num;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}