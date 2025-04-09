import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class FindTheKThCharacterInStringGameITest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.kthCharacter(1)).isEqualTo('a');
        then(solution.kthCharacter(2)).isEqualTo('b');
        then(solution.kthCharacter(3)).isEqualTo('b');
        then(solution.kthCharacter(4)).isEqualTo('c');
        then(solution.kthCharacter(5)).isEqualTo('b');
        then(solution.kthCharacter(10)).isEqualTo('c');
        then(solution.kthCharacter(50)).isEqualTo('d');
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public char kthCharacter(int k) {
            StringBuilder buf = new StringBuilder();
            buf.append('a');
            while (buf.length() < k) {
                int lng = buf.length();
                for (int j = 0; j < lng; j++) {
                    buf.append((char) (buf.charAt(j) + 1));
                }
            }
            return buf.charAt(k - 1);
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}