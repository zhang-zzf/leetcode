import org.junit.jupiter.api.Test;


class TheTwoSneakyNumbersOfDigitvilleTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] getSneakyNumbers(int[] nums) {
            int[] cnt = new int[101];
            int[] ans = new int[2];
            int idx = 0;
            for (int num : nums) {
                cnt[num] += 1;
                if (cnt[num] == 2) {
                    ans[idx++] = num;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}