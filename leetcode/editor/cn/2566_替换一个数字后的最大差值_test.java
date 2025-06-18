import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.assertj.core.api.BDDAssertions.then;


class MaximumDifferenceByRemappingADigitTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minMaxDifference(11891)).isEqualTo(99009);
        then(solution.minMaxDifference(90)).isEqualTo(99);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int minMaxDifference(int num) {
            return version2(num);
        }

        private static int version2(int num) {
            int cnt = 0;
            for (int n = num; n != 0; n /= 10) {
                cnt += 1;
            }
            int[] digits = new int[cnt];
            for (int n = num, c = cnt; n != 0; n /= 10) {
                digits[--c] = n % 10;
            }
            //
            // 最大值
            // 从高位选取第1个非9的数字替换
            int dh = digits[0];
            for (int i = 0; i < digits.length; i++) {
                if (digits[i] != 9) {
                    dh = digits[i];
                    break;
                }
            }
            int max = replace(Arrays.copyOf(digits, digits.length), dh, 9);
            int min = replace(Arrays.copyOf(digits,digits.length), digits[0], 0);
            return max - min;
        }

        private static int replace(int[] digits, int pivot, int target) {
            for (int i = 0; i < digits.length; i++) {
                if (digits[i] == pivot) {
                    digits[i] = target;
                }
            }
            int ans = 0;
            for (int digit : digits) {
                ans = ans * 10 + digit;
            }
            return ans;
        }

        // then(solution.minMaxDifference(90)).isEqualTo(99);
        // 不一定替换第一位
        private static int failedCase1(int num) {
            int cnt = 0;
            for (int n = num; n != 0; n /= 10) {
                cnt += 1;
            }
            int[] digits = new int[cnt];
            for (int n = num, c = cnt; n != 0; n /= 10) {
                digits[--c] = n % 10;
            }
            int dh = digits[0];
            int diff = 0;
            for (int i = 0; i < digits.length; i++) {
                if (digits[i] != dh) {
                    diff = diff * 10 + 0;
                } else {
                    diff = diff * 10 + 9;
                }
            }
            return diff;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}