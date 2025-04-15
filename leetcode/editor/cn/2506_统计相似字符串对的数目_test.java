import java.util.HashMap;
import java.util.Map;
import org.junit.jupiter.api.Test;


class CountPairsOfSimilarStringsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int similarPairs(String[] words) {
            return v2(words);
        }

        private static int v2(String[] words) {
            int ans = 0;
            Map<Integer, Integer> cnt = new HashMap<>();
            for (String word : words) {
                int mask = 0;
                for (char c : word.toCharArray()) {
                    mask |= 1 << (c - 'a');
                }
                Integer c = cnt.getOrDefault(mask, 0);
                ans += c;
                cnt.put(mask, c + 1);
            }
            return ans;
        }

        private static int v1(String[] words) {
            int ans = 0;
            Map<Integer, Integer> cnt = new HashMap<>();
            for (String word : words) {
                int mask = 0;
                for (char c : word.toCharArray()) {
                    mask |= 1 << (c - 'a');
                }
                cnt.put(mask, cnt.getOrDefault(mask, 0) + 1);
            }
            for (Integer c : cnt.values()) {
                ans += c * (c - 1) / 2;
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}