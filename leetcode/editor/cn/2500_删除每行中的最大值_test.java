import static org.assertj.core.api.BDDAssertions.then;

import java.util.Arrays;
import org.junit.jupiter.api.Test;


class DeleteGreatestValueInEachRowTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.deleteGreatestValue(new int[][]{{1, 2, 4}, {3, 3, 1}})).isEqualTo(8);
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int deleteGreatestValue(int[][] grid) {
            int ans = 0;
            for (int[] arr : grid) {
                Arrays.sort(arr);
            }
            for (int i = 0; i < grid[0].length; i++) {
                int max = Integer.MIN_VALUE;
                for (int[] arr : grid) {
                    max = Math.max(max, arr[i]);
                }
                ans += max;
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}