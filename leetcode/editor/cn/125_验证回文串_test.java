//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。 
//
// 字母和数字都属于字母数字字符。 
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
// 
//
// 示例 2： 
//
// 
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
// 
//
// 示例 3： 
//
// 
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 2 * 10⁵ 
// s 仅由可打印的 ASCII 字符组成 
// 
//
// Related Topics 双指针 字符串 
// 👍 665 👎 0


import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class ValidPalindromeTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        String palindrome = "A man, a plan, a canal: Panama";
        then(solution.isPalindrome(palindrome)).isTrue();
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public boolean isPalindrome(String s) {
            boolean ans = true;
            int left = 0, right = s.length() - 1;
            for (; left < right; ) {
                int left_ = s.charAt(left);
                if (!isalnum(left_)) {
                    left += 1;
                    continue;
                }
                int right_ = s.charAt(right);
                if (!isalnum(right_)) {
                    right -= 1;
                    continue;
                }
                if (lower(left_) != lower(right_)) {
                    ans = false;
                    break;
                }
                left += 1;
                right -= 1;
            }
            return ans;
        }

        private int lower(int ascii) {
            if (ascii >= 'A' && ascii <= 'Z') {
                return ascii + ('a' - 'A');
            }
            return ascii;
        }

        private boolean isalnum(int ascii) {
            if (ascii >= '0' && ascii <= '9') {
                return true;
            }
            if (ascii >= 'A' && ascii <= 'Z') {
                return true;
            }
            if (ascii >= 'a' && ascii <= 'z') {
                return true;
            }
            return false;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}