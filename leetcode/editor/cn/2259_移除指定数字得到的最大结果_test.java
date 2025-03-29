import org.junit.jupiter.api.Test;
import static org.assertj.core.api.BDDAssertions.then;


class RemoveDigitFromNumberToMaximizeResultTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.removeDigit("1231", '1')).isEqualTo("231");
        then(solution.removeDigit("1111", '1')).isEqualTo("111");
    }
    
    //leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public String removeDigit(String number, char digit) {
        String ans = "";
        for (int i = 0; i < number.length(); i++) {
            if (number.charAt(i) == digit) {
                String candidate = new StringBuilder(number).deleteCharAt(i).toString();
                if (candidate.compareTo(ans) > 0) {
                    ans = candidate;
                }
            }
        }
        return ans;
    }
}
//leetcode submit region end(Prohibit modification and deletion)

    
}