import java.util.ArrayList;
import java.util.List;
import org.junit.jupiter.api.Test;


class IntersectionOfMultipleArraysTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public List<Integer> intersection(int[][] nums) {
            List<Integer> ans = new ArrayList<>();
            int[] cnt = new int[1001];
            for (int[] num : nums) {
                for (int n : num) {
                    cnt[n] += 1;
                }
            }
            for (int i = 0; i < cnt.length; i++) {
                if (cnt[i] == nums.length) {
                    ans.add(i);
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}