import org.junit.jupiter.api.Test;


class CountHillsAndValleysInAnArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countHillValley(int[] nums) {
            int ans = 0;
            for (int i = 1; i < nums.length - 1; i++) {
                if (nums[i] == nums[i - 1]) {
                    continue;
                }
                int left = 0, right = 0;
                for (int j = i - 1; j >= 0; j--) {
                    if (nums[j] != nums[i]) {
                        left = nums[i] - nums[j];
                        break;
                    }
                }
                if (left == 0) {
                    continue;
                }
                for (int j = i + 1; j < nums.length; j++) {
                    if (nums[j] != nums[i]) {
                        right = nums[j] - nums[i];
                        break;
                    }
                }
                if (left * right < 0) {
                    ans += 1;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}