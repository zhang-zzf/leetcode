import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CheckDistancesBetweenSameLettersTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        int[] distance = {1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
        then(solution.checkDistances("abaccb", distance)).isTrue();
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean checkDistances(String s, int[] distance) {
            int[] cnt = new int[27];
            for (int i = 0; i < s.length(); i++) {
                char c = s.charAt(i);
                if (cnt[c - 'a'] == 0) {
                    cnt[c - 'a'] = i + 1;
                }
                else {
                    if (i - cnt[c - 'a'] != distance[c - 'a']) {
                        return false;
                    }
                }
            }
            return true;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}