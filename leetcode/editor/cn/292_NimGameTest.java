import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;


class NimGameTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.canWinNim(1)).isTrue();
        then(solution.canWinNim(3)).isTrue();
        then(solution.canWinNim(4)).isFalse();
        then(solution.canWinNim(8)).isFalse();
        then(solution.canWinNim(12)).isFalse();
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        public boolean canWinNim(int n) {
            // 思路很简单
            // n = aNum 能不能赢看
            // n-1 / n-2 / n-3 对手会不会输
            int[] dp = new int[3];
            for (int i = 4; i <= n; i++) {
                boolean canWin = false;
                for (int j = 0; j < 3; j++) {
                    if (dp[j] == 1) {
                        canWin = true;
                        break;
                    }
                }
                dp[i % 3] = canWin ? 0 : 1;
            }
            return dp[n % 3] == 0;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}