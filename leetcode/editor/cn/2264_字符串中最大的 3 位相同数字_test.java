import org.junit.jupiter.api.Test;


class Largest3SameDigitNumberInStringTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public String largestGoodInteger(String num) {
            String ans = "";
            for (int i = 0; i <= num.length() - 3; i++) {
                if (num.charAt(i) == num.charAt(i + 1) && num.charAt(i) == num.charAt(i + 2)) {
                    String n = num.substring(i, i + 3);
                    if (n.compareTo(ans) > 0) {
                        ans = n;
                    }
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}