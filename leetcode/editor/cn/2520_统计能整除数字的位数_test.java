import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CountTheDigitsThatDivideANumberTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.countDigits(121)).isEqualTo(2);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countDigits(int num) {
            return solution(num);
        }

        private int solution(int num) {
            int[] digits = new int[10];
            for (int n = num; n > 0; n /= 10) {
                digits[n % 10] += 1;
            }
            // num -> 0
            int ans = 0;
            for (int i = 1; i < 10; i++) {
                if (digits[i] > 0 && num % i == 0) {
                    ans += digits[i];
                }
            }
            return ans;
        }

        private int failedCase1(int num) {
            int[] digits = new int[10];
            while (num > 0) {
                digits[num % 10] += 1;
                num /= 10;
            }
            // num -> 0
            int ans = 0;
            for (int i = 1; i < 10; i++) {
                if (digits[i] > 0 && num % i == 0) {
                    ans += digits[i];
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}