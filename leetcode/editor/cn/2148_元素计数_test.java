// 给你一个整数数组 nums ，统计并返回在 nums 中同时至少具有一个严格较小元素和一个严格较大元素的元素数目。
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [11,7,2,15]
// 输出：2
// 解释：元素 7 ：严格较小元素是元素 2 ，严格较大元素是元素 11 。
// 元素 11 ：严格较小元素是元素 7 ，严格较大元素是元素 15 。
// 总计有 2 个元素都满足在 nums 中同时存在一个严格较小元素和一个严格较大元素。
// 
//
// 示例 2： 
//
// 
// 输入：nums = [-3,3,3,90]
// 输出：2
// 解释：元素 3 ：严格较小元素是元素 -3 ，严格较大元素是元素 90 。
// 由于有两个元素的值为 3 ，总计有 2 个元素都满足在 nums 中同时存在一个严格较小元素和一个严格较大元素。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// -10⁵ <= nums[i] <= 10⁵ 
// 
//
// Related Topics 数组 计数 排序 
// 👍 28 👎 0


import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CountElementsWithStrictlySmallerAndGreaterElementsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.countElements(new int[]{11, 7, 2, 15})).isEqualTo(2);
        then(solution.countElements(new int[]{-3, 3, 3, 5})).isEqualTo(2);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        /**
         * 思路解析：最小值，最大值之间的所有数字均满足条件
         */
        public int countElements(int[] nums) {
            int min = nums[0], max = nums[0];
            for (int num : nums) {
                if (num > max) {
                    max = num;
                }
                else if (num < min) {
                    min = num;
                }
            }
            int ans = 0;
            for (int num : nums) {
                if (num > min && num < max) {
                    ans += 1;
                }
            }
            return ans;
        }
    }
// leetcode submit region end(Prohibit modification and deletion)


}