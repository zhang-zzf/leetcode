import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class SplitWithMinimumSumTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        int ans = solution.splitNum(4325);
        then(ans).isEqualTo(59);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int splitNum(int num) {
            int[] digits = new int[10];
            while (num > 0) {
                digits[num % 10]++;
                num /= 10;
            }
            int n1 = 0, n2 = 0;
            while (true) {
                int n1_ = min(digits);
                if (n1_ == -1) {
                    break;
                }
                n1 = n1 * 10 + n1_;
                int n2_ = min(digits);
                if (n2_ == -1) {
                    break;
                }
                n2 = n2 * 10 + n2_;
            }
            return n1 + n2;
        }

        int min(int[] digits) {
            int ans = -1;
            for (int i = 0; i < 10; ) {
                if (digits[i] == 0) {
                    i += 1;
                }
                else {
                    digits[i] -= 1;
                    ans = i;
                    break;
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}