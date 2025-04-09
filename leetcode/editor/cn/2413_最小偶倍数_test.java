import org.junit.jupiter.api.Test;


class SmallestEvenMultipleTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int smallestEvenMultiple(int n) {
            if (n % 2 == 0) {
                return n;
            }
            return 2 * n;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}