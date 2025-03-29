import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CheckIfMatrixIsXMatrixTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.checkXMatrix(new int[][]{
            {2, 0, 0, 1},
            {0, 3, 1, 0},
            {0, 5, 2, 0},
            {4, 0, 0, 2}
        })).isTrue();
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean checkXMatrix(int[][] grid) {
            boolean ans = true;
            for (int i = 0; i < grid.length; i++) {
                for (int j = 0; j < grid.length; j++) {
                    // 对角线判断逻辑错误j
                    // if ((i == j && grid[i][j] == 0) || (i != j && grid[i][j] != 0)) {
                    //     ans = false;
                    //     break;
                    // }
                    if (i == j || i + j == grid.length - 1) {// 对角线
                        if (grid[i][j] == 0) {
                            ans = false;
                            break;
                        }
                    }
                    else{
                        if (grid[i][j] != 0) {
                            ans = false;
                            break;
                        }
                    }
                }
            }
            return ans;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}