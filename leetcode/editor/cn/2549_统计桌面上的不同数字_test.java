import org.junit.jupiter.api.Test;


class CountDistinctNumbersOnBoardTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int distinctIntegers(int n) {
            return n == 1 ? 1 : n - 1;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}