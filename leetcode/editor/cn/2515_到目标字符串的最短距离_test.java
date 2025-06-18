import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class ShortestDistanceToTargetStringInACircularArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        String[] words = {"hello", "i", "am", "leetcode", "hello"};
        then(solution.closestTarget(words, "hello", 1)).isEqualTo(1);
        then(solution.closestTarget(words, "hello", 0)).isEqualTo(0);
        then(solution.closestTarget(words, "hello1", 0)).isEqualTo(-1);
        then(solution.closestTarget(words, "hello", 3)).isEqualTo(1);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int closestTarget(String[] words, String target, int startIndex) {
            int length = words.length;
            int step = 0;
            boolean match = false;
            for (int left = startIndex, right = startIndex; step < length; ) {
                if (words[left].equals(target) || words[right].equals(target)) {
                    match = true;
                    break;
                }
                left = (left - 1 + length) % length;
                right = (right + 1) % length;
                step += 1;
            }
            return match ? step : -1;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}