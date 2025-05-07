import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class MaximumCountOfPositiveIntegerAndNegativeIntegerTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.maximumCount(new int[]{-2, -1, -1, 1, 2, 3})).isEqualTo(3);
        then(solution.maximumCount(new int[]{-3, -2, -1, 0, 0, 1, 2})).isEqualTo(3);
        then(solution.maximumCount(new int[]{5, 20, 66, 1314})).isEqualTo(4);
        // failed case
        then(solution.maximumCount(new int[]{0, 0})).isEqualTo(0);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int maximumCount(int[] nums) {
            // 找到 <=0 的最右边位置
            int left = 0, right = nums.length - 1;
            // -1 0 0 1
            // -1 -1 -1
            // 1 1 1
            // 0 1 1
            while (left <= right) {
                int mid = left + (right - left) / 2;
                if (nums[mid] <= 0) {
                    left = mid + 1;
                }
                else {
                    right = mid - 1;
                }
            }
            // left 指向第一个 > 0 的位置
            int posCnt = nums.length - left;
            do {
                left -= 1;
            } while (left >= 0 && nums[left] >= 0);
            int negCnt = left + 1;
            return Math.max(posCnt, negCnt);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}