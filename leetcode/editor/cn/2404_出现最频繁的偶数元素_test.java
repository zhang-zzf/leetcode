import java.util.HashMap;
import java.util.Map;
import org.junit.jupiter.api.Test;


class MostFrequentEvenElementTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int mostFrequentEven(int[] nums) {
            Map<Integer, Integer> map = new HashMap<>();
            int ans = -1;
            int maxCnt = 0;
            for (int num : nums) {
                if (num % 2 == 1) {
                    continue;
                }
                int cnt = map.getOrDefault(num, 0) + 1;
                map.put(num, cnt);
                maxCnt = Math.max(maxCnt, cnt);
            }
            if (map.isEmpty()) {
                return ans;
            }
            ans = Integer.MAX_VALUE;
            for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
                if (entry.getValue() == maxCnt) {
                    ans = Math.min(ans, entry.getKey());
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}