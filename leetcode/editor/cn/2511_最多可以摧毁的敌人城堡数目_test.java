import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class MaximumEnemyFortsThatCanBeCapturedTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.captureForts(new int[]{1, 0, 0, -1, 0, 0, 0, 0, 1})).isEqualTo(4);
    }

    @Test
    void givenFailedCase1_when_thenSuccess() {
        then(solution.captureForts(new int[]{0, 0, 1, -1})).isEqualTo(0);
    }

    @Test
    void givenFailedCase2_when_thenSuccess() {
        then(solution.captureForts(new int[]{-1, -1, 0, 1, 0, 0, 1, -1, 1, 0})).isEqualTo(1);
    }


    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int captureForts(int[] forts) {
            return v3(forts);
        }

        private static int v3(int[] forts) {
            int ans = 0;
            for (int left = 0; left < forts.length; ) {
                if (forts[left] == 0) {
                    left += 1;
                    continue;
                }
                int right = left + 1;
                for (; right < forts.length; right++) {
                    if (forts[right] != 0) {
                        break;
                    }
                }
                if (right >= forts.length) {// 结束
                    break;
                }
                if (forts[left] + forts[right] == 0) {// -1+1 or 1+-1
                    ans = Math.max(ans, right - left - 1);
                }
                left = right;
            }
            return ans;
        }

        /**
         * 题目理解错误：将你的军队从某个你控制的城堡位置 i 移动到一个空的位置 j 必须有一个位置为空堡垒
         */
        private static int v2(int[] forts) {
            int ans = 0;
            for (int left = 0; left < forts.length; ) {
                if (forts[left] == 0) {
                    left += 1;
                    continue;
                }
                int right = left + 1;
                for (; right < forts.length; right++) {
                    if (forts[right] != 0) {
                        break;
                    }
                }
                if (right >= forts.length) {// 结束
                    break;
                }
                if (forts[left] == 1 || forts[right] == 1) {
                    ans = Math.max(ans, right - left - 1);
                }
                left = right;
            }
            return ans;
        }

        private static int failedCase1(int[] forts) {
            int ans = 0;
            for (int left = 0; left < forts.length; left++) {
                if (forts[left] == 0) {
                    continue;
                }
                int right = left + 1;
                for (; right < forts.length; right++) {
                    if (forts[right] != 0) {
                        break;
                    }
                }
                if (forts[left] == 1 || forts[right] == 1) {
                    ans = Math.max(ans, right - left - 1);
                }
            }
            return ans;
        }
    }

    // leetcode submit region end(Prohibit modification and deletion)
}