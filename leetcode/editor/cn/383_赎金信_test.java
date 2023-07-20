//给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。 
//
// 如果可以，返回 true ；否则返回 false 。 
//
// magazine 中的每个字符只能在 ransomNote 中使用一次。 
//
// 
//
// 示例 1： 
//
// 
//输入：ransomNote = "a", magazine = "b"
//输出：false
// 
//
// 示例 2： 
//
// 
//输入：ransomNote = "aa", magazine = "ab"
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：ransomNote = "aa", magazine = "aab"
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 1 <= ransomNote.length, magazine.length <= 10⁵ 
// ransomNote 和 magazine 由小写英文字母组成 
// 
//
// Related Topics 哈希表 字符串 计数 
// 👍 744 👎 0


import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class RansomNoteTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        boolean ans = solution.canConstruct("aa", "aba");
        then(ans).isTrue();
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public boolean canConstruct(String ransomNote, String magazine) {
            boolean ans = true;
            int[] mapping = new int[26];
            for (int c : magazine.toCharArray()) {
                mapping[c - 'a'] += 1;
            }
            for (int c : ransomNote.toCharArray()) {
                int cnt = mapping[c - 'a'];
                if (cnt < 1) {
                    ans = false;
                    break;
                }
                mapping[c - 'a'] -= 1;
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}