import static org.assertj.core.api.BDDAssertions.then;

import java.util.HashMap;
import java.util.Map;
import org.junit.jupiter.api.Test;


class DecodeTheMessageTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
            .isEqualTo("this is a secret");
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String decodeMessage(String key, String message) {
            StringBuilder ans = new StringBuilder();
            char[] mapping = new char[26];
            char curChar = 'a';
            for (char c : key.toCharArray()) {
                if (c == ' ' || mapping[c - 'a'] > 0) {
                    continue;
                }
                mapping[c - 'a'] = curChar++;
            }
            for (char c : message.toCharArray()) {
                if (c == ' ') {
                    ans.append(' ');
                }
                else {
                    ans.append(mapping[c - 'a']);
                }
            }
            return ans.toString();
        }

        private static void v1(String key, String message, StringBuilder ans) {
            Map<Character, Character> mapping = new HashMap<>();
            char curChar = 'a';
            for (char c : key.toCharArray()) {
                if (c == ' ' || mapping.containsKey(c)) {
                    continue;
                }
                mapping.put(c, curChar++);
            }
            for (char c : message.toCharArray()) {
                ans.append(mapping.getOrDefault(c, ' '));
            }
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}