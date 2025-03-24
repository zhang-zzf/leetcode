import org.junit.jupiter.api.Test;


class CountIntegersWithEvenDigitSumTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countEven(int num) {
            int ans = 0;
            for (int i = 1; i <= num; i++) {
                if (sumDigit(i) % 2 == 0) {
                    ans += 1;
                }
            }
            return ans;
        }

        private int sumDigit(int i) {
            int sum = 0;
            for (; i > 0; i /= 10) {
                sum = sum + i % 10;
            }
            return sum;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}