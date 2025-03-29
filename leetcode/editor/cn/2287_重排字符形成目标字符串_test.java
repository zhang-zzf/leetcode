import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class RearrangeCharactersToMakeTargetStringTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.rearrangeCharacters("ilovecodingonleetcode", "code"))
            .isEqualTo(2);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public static final int CHAR_CNT = 26;

        public int rearrangeCharacters(String s, String target) {
            int ans = Integer.MAX_VALUE;
            int[] sCnt = new int[CHAR_CNT];
            int[] tCnt = new int[CHAR_CNT];
            for (char c : s.toCharArray()) {
                sCnt[c - 'a'] += 1;
            }
            for (char c : target.toCharArray()) {
                tCnt[c - 'a'] += 1;
            }
            for (int c = 0; c < CHAR_CNT; c++) {
                if (tCnt[c] == 0) {
                    continue;
                }
                ans = Math.min(sCnt[c] / tCnt[c], ans);
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}