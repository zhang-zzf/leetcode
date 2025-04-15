import org.junit.jupiter.api.Test;


class FindThePivotIntegerTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int pivotInteger(int n) {
            int sum = 0, prefixSum = 0;
            for (int i = 1; i <= n; i++) {
                sum += i;
            }
            for (int i = 1; i <= n; i++) {
                prefixSum += i;
                if (prefixSum == sum - prefixSum + i) {
                    return i;
                }
            }
            return -1;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}