import org.junit.jupiter.api.Test;


class DivideArrayIntoEqualPairsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean divideArray(int[] nums) {
            boolean ans = true;
            int lng = 501;
            int[] cnt = new int[lng];
            for (int num : nums) {
                cnt[num] += 1;
            }
            for (int i = 0; i < lng; i++) {
                if (cnt[i] % 2 != 0) {
                    ans = false;
                    break;
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}