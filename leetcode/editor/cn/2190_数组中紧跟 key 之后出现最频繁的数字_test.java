import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;
import org.junit.jupiter.api.Test;


class MostFrequentNumberFollowingKeyInAnArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int mostFrequent(int[] nums, int key) {
            int ans = 0;
            Map<Integer, Integer> mapping = new HashMap<>();
            for (int i = 0; i < nums.length - 1; i++) {
                if (nums[i] == key) {
                    mapping.put(nums[i + 1], mapping.getOrDefault(nums[i + 1], 0) + 1);
                }
            }
            for (java.util.Map.Entry<Integer, Integer> entry : mapping.entrySet()) {
                if (entry.getValue() > mapping.getOrDefault(ans, 0)) {
                    ans = entry.getKey();
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}