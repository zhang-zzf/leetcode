import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class LeftAndRightSumDifferencesTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        int[] ans = solution.leftRightDifference(new int[]{10, 4, 8, 3});
        then(ans).isEqualTo(new int[]{15, 1, 11, 22});
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] leftRightDifference(int[] nums) {
            int sum = 0;
            for (int num : nums) {
                sum += num;
            }
            int leftSum = 0;
            int rightSum = sum;
            int[] ans = new int[nums.length];
            for (int i = 0; i < ans.length; i++) {
                rightSum -= nums[i];
                ans[i] = Math.abs(leftSum - rightSum);
                leftSum += nums[i];
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}