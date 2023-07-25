//给定两个字符串 s 和 t ，判断它们是否是同构的。 
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。 
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。 
//
// 
//
// 示例 1: 
//
// 
//输入：s = "egg", t = "add"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "foo", t = "bar"
//输出：false 
//
// 示例 3： 
//
// 
//输入：s = "paper", t = "title"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 
//
// 
// 1 <= s.length <= 5 * 10⁴ 
// t.length == s.length 
// s 和 t 由任意有效的 ASCII 字符组成 
// 
//
// Related Topics 哈希表 字符串 
// 👍 621 👎 0


import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class IsomorphicStringsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        boolean ans = solution.isIsomorphic("egg", "add");
        then(ans).isTrue();
    }

    /**
     * "badc"
     * "baba"
     */
    @Test
    void givenFailedCase1_when_thenSuccess() {
        boolean ans = solution.isIsomorphic("badc", "baba");
        then(ans).isFalse();
    }


    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public static final int NO_SET = -1;
        public static final int ASCII_LNG = 128;

        public boolean isIsomorphic(String s, String t) {
            boolean ans = true;
            int[] m1 = new int[ASCII_LNG];
            int[] m2 = new int[ASCII_LNG];
            for (int i = 0; i < ASCII_LNG; i++) {
                m1[i] = NO_SET;
                m2[i] = NO_SET;
            }
            for (int i = 0; i < s.length(); i++) {
                int sc = s.charAt(i);
                char tc = t.charAt(i);
                if (m1[sc] == NO_SET) {
                    if (m2[tc] == NO_SET || m2[tc] == sc) {
                        m1[sc] = tc;
                        m2[tc] = sc;
                    } else {
                        ans = false;
                        break;
                    }
                } else if (m1[sc] != tc) {
                    ans = false;
                    break;
                }
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}