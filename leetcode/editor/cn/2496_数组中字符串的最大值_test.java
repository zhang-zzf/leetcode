import org.junit.jupiter.api.Test;


class MaximumValueOfAStringInAnArrayTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int maximumValue(String[] strs) {
            int ans = 0;
            for (String str : strs) {
                boolean isNum = true;
                for (char c : str.toCharArray()) {
                    if (!Character.isDigit(c)) {
                        isNum = false;
                        break;
                    }
                }
                if (isNum) {
                    ans = Math.max(ans, Integer.parseInt(str));
                }
                else {
                    ans = Math.max(ans, str.length());
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}