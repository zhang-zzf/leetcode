import org.junit.jupiter.api.Test;


class MaximumNumberOfPairsInArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] numberOfPairs(int[] nums) {
            int[] ans = new int[2];
            int[] cnt = new int[101];
            for (int num : nums) {
                cnt[num] += 1;
            }
            for (int j : cnt) {
                ans[0] += j / 2;
                ans[1] += j % 2;
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)


}