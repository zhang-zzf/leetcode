import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CircularSentenceTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.isCircularSentence("leetcode exercises sound delightful")).isTrue();
        then(solution.isCircularSentence("eetcode")).isTrue();
        then(solution.isCircularSentence("Leetcode is cool")).isFalse();
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean isCircularSentence(String sentence) {
            boolean ans = true;
            String[] words = sentence.split(" ");
            for (int i = 0; i < words.length; i++) {
                int j = (i + 1) % words.length;
                if (words[i].charAt(words[i].length() - 1) != words[j].charAt(0)) {
                   ans = false;
                   break;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}