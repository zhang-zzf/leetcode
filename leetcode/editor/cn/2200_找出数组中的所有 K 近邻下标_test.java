import java.util.ArrayList;
import java.util.List;
import org.junit.jupiter.api.Test;


class FindAllKDistantIndicesInAnArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public List<Integer> findKDistantIndices(int[] nums, int key, int k) {
            int[] flag = new int[nums.length];
            for (int i = 0; i < nums.length; i++) {
                if (nums[i] == key) {
                    for (int j = i - k; j <= i + k; j++) {
                        if (j >= 0 && j < nums.length) {
                            flag[j] = 1;
                        }
                    }
                }
            }
            List<Integer> ans = new ArrayList<>();
            for (int i = 0; i < nums.length; i++) {
                if (flag[i] == 1) {
                    ans.add(i);
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}