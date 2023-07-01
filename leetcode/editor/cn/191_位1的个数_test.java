import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class NumberOf1BitsTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.hammingWeight(3)).isEqualTo(2);
        then(solution.hammingWeight(-1)).isEqualTo(32);
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    public class Solution {

        // you need to treat n as an unsigned value
        public int hammingWeight(int n) {
            int ans = 0;
            while (n != 0) {
                n &= (n - 1);
                ans += 1;
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}