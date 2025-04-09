import static org.assertj.core.api.BDDAssertions.then;

import java.util.Arrays;
import org.junit.jupiter.api.Test;


class SortThePeopleTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        String[] names = {"Mary", "John", "Emma"};
        int[] heights = {180, 165, 170};
        then(solution.sortPeople(names, heights)).containsExactly("Mary", "Emma", "John");
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String[] sortPeople(String[] names, int[] heights) {
            String[] ans = new String[names.length];
            Integer[] indices = new Integer[names.length];
            for (int i = 0; i < names.length; i++) {
                indices[i] = i;
            }
            Arrays.sort(indices, (i, j) -> heights[j] - heights[i]);
            for (int i = 0; i < names.length; i++) {
                ans[i] = names[indices[i]];
            }
            return ans;
        }

        private static String[] v1(String[] names, int[] heights) {
            String[] ans = new String[names.length];
            int lastMax = Integer.MAX_VALUE;
            for (int i = 0; i < names.length; i++) {
                int curMax = 0, index = 0;
                for (int j = 0; j < names.length; j++) {
                    if (heights[j] < lastMax && heights[j] > curMax) {
                        curMax = heights[j];
                        index = j;
                    }
                }
                lastMax = curMax;
                ans[i] = names[index];
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}