import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class NumberOfArithmeticTripletsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.arithmeticTriplets(new int[]{0, 1, 4, 6, 7, 10}, 3))
                .isEqualTo(2);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int arithmeticTriplets(int[] nums, int diff) {
            int[] reverse = new int[202];
            for (int i = 0; i < nums.length; i++) {
                reverse[nums[i]] = i + 1; // + 1 很关键
            }
            int ans = 0;
            for (int num = 0; num < reverse.length; num++) {
                int i = reverse[num], j = 0, k = 0;
                if (i == 0) {
                    continue;
                }
                if ((num + diff < reverse.length && (j = reverse[num + diff]) > 0)
                        && (num + diff * 2 < reverse.length && (k = reverse[num + diff * 2]) > 0)) {
                    if (i < j && j < k) {
                        ans += 1;
                    }
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}