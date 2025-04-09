import org.junit.jupiter.api.Test;

import java.util.HashMap;
import java.util.Map;


class FindSubarraysWithEqualSumTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean findSubarrays(int[] nums) {
            Map<Integer, Integer> map = new HashMap<>();
            for (int i = 0; i < nums.length - 1; i++) {
                int sum = nums[i] + nums[i + 1];
                if (map.containsKey(sum)) {
                    return true;
                }
                map.put(sum, 1);
            }
            return false;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}