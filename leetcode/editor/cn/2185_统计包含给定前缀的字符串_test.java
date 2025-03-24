import org.junit.jupiter.api.Test;


class CountingWordsWithAGivenPrefixTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int prefixCount(String[] words, String pref) {
            int ans = 0;
            for (String word : words) {
                if (word.startsWith(pref)) {
                    ans += 1;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}