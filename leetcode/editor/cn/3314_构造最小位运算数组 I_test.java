import static org.assertj.core.api.BDDAssertions.then;

import java.util.List;
import org.junit.jupiter.api.Test;


class ConstructTheMinimumBitwiseArrayITest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.minBitwiseArray(List.of(11, 13, 31)))
            .containsExactly(9, 12, 15);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] minBitwiseArray(List<Integer> nums) {
            int[] ans = new int[nums.size()];
            for (int i = 0; i < nums.size(); i++) {
                Integer num = nums.get(i);
                ans[i] = -1;
                for (int j = 0; j < num; j++) {
                    if ((j | j + 1) == num) {
                        ans[i] = j;
                        break;
                    }
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}