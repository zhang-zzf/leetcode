import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class NimGameTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.canWinNim(1)).isTrue();
        then(solution.canWinNim(4)).isFalse();
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public boolean canWinNim(int n) {
            return (n % 4) != 0;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}