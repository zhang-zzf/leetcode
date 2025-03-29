import org.junit.jupiter.api.Test;


class CountPrefixesOfAGivenStringTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countPrefixes(String[] words, String s) {
            int ans = 0;
            for (String word : words) {
                if (s.startsWith(word)) {
                    ans += 1;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}