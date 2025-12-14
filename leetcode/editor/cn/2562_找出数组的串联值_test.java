import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class FindTheArrayConcatenationValueTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.findTheArrayConcVal(new int[]{7, 52, 2, 4})).isEqualTo(596);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public long findTheArrayConcVal(int[] nums) {
            long ans = 0;
            for (int left = 0, right = nums.length - 1; left <= right; left++, right--) {
                if (left == right) {
                    ans += nums[left];
                }
                else {
                    ans += Long.parseLong(nums[left] + "" + nums[right]);
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}