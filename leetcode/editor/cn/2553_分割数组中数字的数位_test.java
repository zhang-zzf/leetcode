import org.junit.jupiter.api.Test;


class SeparateTheDigitsInAnArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] separateDigits(int[] nums) {
            int n = 0;
            String[] numStr = new String[nums.length];
            for (int i = 0; i < nums.length; i++) {
                numStr[i] = String.valueOf(nums[i]);
                n += numStr[i].length();
            }
            int[] ret = new int[n];
            int idx = 0;
            for (String s : numStr) {
                for (char c : s.toCharArray()) {
                    ret[idx++] = c - '0';
                }
            }
            return ret;
        }


    }
    // leetcode submit region end(Prohibit modification and deletion)


}