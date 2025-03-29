import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CalculateAmountPaidInTaxesTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.calculateTax(new int[][]{{3, 50}, {7, 10}, {12, 25}}, 10))
            .isEqualTo(2.65);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        // 最后一个税级的上限大于等于 income
        public double calculateTax(int[][] brackets, int income) {
            double ans = 0;
            for (int i = 0; i < brackets.length; i++) {
                int cur = brackets[i][0];
                int toTax = Math.min(cur, income) - (i == 0 ? 0 : brackets[i - 1][0]);
                ans += toTax * brackets[i][1] / 100.0;
                if (income <= cur) {
                    break;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}