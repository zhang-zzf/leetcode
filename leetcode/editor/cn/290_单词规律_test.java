import org.junit.jupiter.api.Test;

import java.util.HashMap;
import java.util.Map;

import static org.assertj.core.api.BDDAssertions.then;


class WordPatternTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.wordPattern("abba", "d a a d")).isTrue();
        then(solution.wordPattern("abba", "d d d d")).isFalse();
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public boolean wordPattern(String pattern, String s) {
            String[] words = s.split(" ");
            if (pattern.length() != words.length) {
                return false;
            }
            boolean ans = true;
            Map<Integer, String> p2w = new HashMap<>();
            Map<String, Integer> w2p = new HashMap<>();
            for (int i = 0; i < words.length; i++) {
                Integer p = (int) pattern.charAt(i);
                String w = words[i];
                String m1 = p2w.get(p);
                Integer m2 = w2p.get(w);
                if (m1 == null && m2 == null) {
                    p2w.put(p, w);
                    w2p.put(w, p);
                } else if (!p.equals(m2) || !w.equals(m1)) {
                    ans = false;
                    break;
                }
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}