import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class MinMaxGameTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minMaxGame(new int[]{1, 3, 5, 2, 4, 8, 2, 2})).isEqualTo(1);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int minMaxGame(int[] nums) {
            return game(nums)[0];
        }

        public int[] game(int[] nums) {
            if (nums.length == 1) {
                return nums;
            }
            int[] newNums = new int[nums.length / 2];
            for (int i = 0; i < nums.length / 2; i++) {
                if (i % 2 == 0) {
                    newNums[i] = Math.min(nums[2 * i], nums[2 * i + 1]);
                } else {
                    newNums[i] = Math.max(nums[2 * i], nums[2 * i + 1]);
                }
            }
            return game(newNums);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}