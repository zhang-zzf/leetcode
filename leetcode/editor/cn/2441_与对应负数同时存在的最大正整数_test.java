import java.util.HashSet;
import java.util.Set;
import org.junit.jupiter.api.Test;


class LargestPositiveIntegerThatExistsWithItsNegativeTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int findMaxK(int[] nums) {
            Set<Integer> set = new HashSet<>();
            for (int num : nums) {
                set.add(num);
            }
            int ans = -1;
            for (int num : nums) {
                if (num > ans && set.contains(-num)) {
                    ans = num;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}