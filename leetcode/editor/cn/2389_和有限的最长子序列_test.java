import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.assertj.core.api.BDDAssertions.then;


class LongestSubsequenceWithLimitedSumTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.answerQueries(new int[]{4, 5, 2, 1}, new int[]{3, 10, 21}))
                .containsExactly(2, 3, 4);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] answerQueries(int[] nums, int[] queries) {
            return failedCase1(nums, queries);
        }

        private int[] failedCase1(int[] nums, int[] queries) {
            int[] cnt = new int[nums.length + 1];
            Arrays.sort(nums);
            for (int i = 0; i < nums.length; i++) {
                cnt[i + 1] = cnt[i] + nums[i];
            }
            int[] ans = new int[queries.length];
            for (int i = 0; i < queries.length; i++) {
                int idx = Arrays.binarySearch(cnt, queries[i]);
                if (idx < 0) {
                    // idx 是插入的位置
                    idx = -idx - 1;
                    // 题目要求比 queries[i] 小的元素的数量
                    ans[i] = idx - 1;
                } else {
                    ans[i] = idx;
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}