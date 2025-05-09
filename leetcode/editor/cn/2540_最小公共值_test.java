import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class MinimumCommonValueTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.getCommon(new int[]{1, 2, 3}, new int[]{2, 4})).isEqualTo(2);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int getCommon(int[] nums1, int[] nums2) {
            int ans = -1;
            int up = 0, down = 0;
            while (up < nums1.length && down < nums2.length) {
                if (nums1[up] == nums2[down]) {
                    ans = nums1[up];
                    break;
                }
                if (nums1[up] < nums2[down]) {
                    up++;
                }
                else {
                    down++;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}