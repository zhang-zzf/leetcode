// 给你一个整数数组 nums ，另给你一个整数 original ，这是需要在 nums 中搜索的第一个数字。
//
// 接下来，你需要按下述步骤操作： 
//
// 
// 如果在 nums 中找到 original ，将 original 乘以 2 ，得到新 original（即，令 original = 2 * 
// original）。
// 否则，停止这一过程。 
// 只要能在数组中找到新 original ，就对新 original 继续 重复 这一过程。 
// 
//
// 返回 original 的 最终 值。 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [5,3,6,1,12], original = 3
// 输出：24
// 解释：
//- 3 能在 nums 中找到。3 * 2 = 6 。
//- 6 能在 nums 中找到。6 * 2 = 12 。
//- 12 能在 nums 中找到。12 * 2 = 24 。
//- 24 不能在 nums 中找到。因此，返回 24 。
// 
//
// 示例 2： 
//
// 
// 输入：nums = [2,7,9], original = 4
// 输出：4
// 解释：
//- 4 不能在 nums 中找到。因此，返回 4 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 1000 
// 1 <= nums[i], original <= 1000 
// 
//
// Related Topics 数组 哈希表 排序 模拟 
// 👍 25 👎 0


import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class KeepMultiplyingFoundValuesByTwoTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        int ans = solution.findFinalValue(new int[]{5, 3, 6, 1, 12}, 3);
        then(ans).isEqualTo(24);
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        int ans = solution.findFinalValue(
            new int[]{161, 28, 640, 264, 81, 561, 320, 2, 61, 244, 183, 108, 773, 61, 976, 122, 988, 2, 370, 392, 488,
                375, 349, 432, 713, 563}, 61);
        then(ans).isEqualTo(1952);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        /**
         * <pre>
         *     1 <= nums.length <= 1000
         *     1 <= nums[i], original <= 1000
         * </pre>
         */
        public int findFinalValue(int[] nums, int original) {
            int[] counter = new int[1001];
            for (int num : nums) {
                counter[num] += 1;
            }
            int ans = original;
            while (ans < 1001 && counter[ans] > 0) { // 边界条件
                ans *= 2;
            }
            return ans;
        }

        private int failedCase1(int[] nums, int original) {
            int[] counter = new int[1001];
            for (int num : nums) {
                counter[num] += 1;
            }
            int ans = original;
            while (counter[ans] > 0) { // 忽略了边界条件
                ans *= 2;
            }
            return ans;
        }
    }
// leetcode submit region end(Prohibit modification and deletion)


}