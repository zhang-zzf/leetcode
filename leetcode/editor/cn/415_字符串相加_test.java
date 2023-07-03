//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。 
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。 
//
// 
//
// 示例 1： 
//
// 
//输入：num1 = "11", num2 = "123"
//输出："134"
// 
//
// 示例 2： 
//
// 
//输入：num1 = "456", num2 = "77"
//输出："533"
// 
//
// 示例 3： 
//
// 
//输入：num1 = "0", num2 = "0"
//输出："0"
// 
//
// 
//
// 
//
// 提示： 
//
// 
// 1 <= num1.length, num2.length <= 10⁴ 
// num1 和num2 都只包含数字 0-9 
// num1 和num2 都不包含任何前导零 
// 
//
// Related Topics 数学 字符串 模拟 
// 👍 725 👎 0


import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class AddStringsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.addStrings("11", "123")).isEqualTo("134");
        then(solution.addStrings("0", "0")).isEqualTo("0");
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.addStrings("1", "9")).isEqualTo("10");
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public String addStrings(String num1, String num2) {
            int lng2 = num2.length();
            int lng1 = num1.length();
            StringBuilder sb = new StringBuilder();
            int promotion = 0;
            for (int i1 = lng1 - 1, i2 = lng2 - 1; i1 >= 0 || i2 >= 0; i1--, i2--) {
                int n1 = 0, n2 = 0;
                if (i1 >= 0) {
                    n1 = num1.charAt(i1) - '0';
                }
                if (i2 >= 0) {
                    n2 = num2.charAt(i2) - '0';
                }
                int sum = n1 + n2 + promotion;
                if (sum < 10) {
                    promotion = 0;
                } else {
                    promotion = 1;
                    sum = sum - 10;
                }
                // char + int => int
                sb.append((char) ('0' + sum));
            }
            if (promotion == 1) {
                sb.append((char) ('0' + promotion));
            }
            return sb.reverse().toString();
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}