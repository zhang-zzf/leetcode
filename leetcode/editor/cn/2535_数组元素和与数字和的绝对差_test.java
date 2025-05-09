import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class DifferenceBetweenElementSumAndDigitSumOfAnArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.differenceOfSum(new int[]{1, 15, 6, 3})).isEqualTo(9);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int differenceOfSum(int[] nums) {
            int sum = 0, digitSum = 0;
            for (int num : nums) {
                sum += num;
                while (num > 0) {
                    digitSum += num % 10;
                    num /= 10;
                }
            }
            return Math.abs(sum - digitSum);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}