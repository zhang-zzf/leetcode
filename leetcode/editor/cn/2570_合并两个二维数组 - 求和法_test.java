import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class MergeTwo2dArraysBySummingValuesTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.mergeArrays(new int[][]{{1, 2}, {2, 3}, {4, 5}}, new int[][]{{1, 4}, {3, 2}, {4, 1}}))
                .isEqualTo(new int[][]{{1, 6}, {2, 3}, {3, 2}, {4, 6}});
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[][] mergeArrays(int[][] nums1, int[][] nums2) {
            int[] arr = new int[1001];
            for (int[] ints : nums1) {
                arr[ints[0]] += ints[1];
            }
            for (int[] ints : nums2) {
                arr[ints[0]] += ints[1];
            }
            int cnt = 0;
            for (int i : arr) {
                if (i > 0) {
                    cnt += 1;
                }
            }
            int[][] ans = new int[cnt][2];
            for (int j = 1000; j >= 0; j--) {
                if (arr[j] > 0) {
                    ans[--cnt] = new int[]{j, arr[j]};
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}