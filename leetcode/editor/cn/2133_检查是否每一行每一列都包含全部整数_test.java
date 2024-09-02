//对一个大小为 n x n 的矩阵而言，如果其每一行和每一列都包含从 1 到 n 的 全部 整数（含 1 和 n），则认为该矩阵是一个 有效 矩阵。 
//
// 给你一个大小为 n x n 的整数矩阵 matrix ，请你判断矩阵是否为一个有效矩阵：如果是，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：matrix = [[1,2,3],[3,1,2],[2,3,1]]
//输出：true
//解释：在此例中，n = 3 ，每一行和每一列都包含数字 1、2、3 。
//因此，返回 true 。
// 
//
// 示例 2： 
//
// 
//
// 
//输入：matrix = [[1,1,1],[1,2,3],[1,2,3]]
//输出：false
//解释：在此例中，n = 3 ，但第一行和第一列不包含数字 2 和 3 。
//因此，返回 false 。
// 
//
// 
//
// 提示： 
//
// 
// n == matrix.length == matrix[i].length 
// 1 <= n <= 100 
// 1 <= matrix[i][j] <= n 
// 
//
// Related Topics 数组 哈希表 矩阵 
// 👍 26 👎 0


import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CheckIfEveryRowAndColumnContainsAllNumbersTest {

  final Solution solution = new Solution();

  @Test
  void givenNormal_when_thenSuccess() {
    boolean ans = solution.checkValid(new int[][]{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}});
    then(ans).isTrue();
  }

  //leetcode submit region begin(Prohibit modification and deletion)
  class Solution {
    public boolean checkValid(int[][] matrix) {
      int lng = matrix.length;
      for (int i = 0; i < lng; i++) {
        int[] row = new int[lng + 1], col = new int[lng + 1];
        for (int j = 0; j < lng; j++) {
          if (row[matrix[i][j]] > 0 || col[matrix[j][i]] > 0) {
            return false;
          }
          row[matrix[i][j]] += 1;
          col[matrix[j][i]] += 1;
        }
      }
      return true;
    }

    private boolean success1(int[][] matrix) {
      int lng = matrix.length;
      for (int i = 0; i < lng; i++) {
        int[] mark = new int[lng];
        for (int j = 0; j < lng; j++) {
          int num_1 = matrix[i][j] - 1;
          if (mark[num_1] > 0) {
            return false;
          }
          mark[num_1] = 1;
        }
        mark = new int[lng];
        for (int j = 0; j < lng; j++) {
          int num_1 = matrix[j][i] - 1;
          if (mark[num_1] > 0) {
            return false;
          }
          mark[num_1] = 1;
        }
      }
      return true;
    }

    /**
     * [[3,1,1],[2,3,2],[2,1,3]]
     */
    private boolean failedCase3(int[][] matrix) {
      int lng = matrix.length;
      for (int i = 0; i < lng; i++) {
        int[] mark = new int[lng];
        for (int j = 0; j < lng; j++) {
          mark[matrix[i][j] - 1] += 1;
        }
        for (int j = 0; j < lng; j++) {
          mark[matrix[j][i] - 1] += 1;
        }
        for (int j = 0; j < lng; j++) {
          if (mark[j] != 2) {
            return false;
          }
        }
      }
      return true;
    }

    /**
     * {{2,2,2},{2,2,2},{2,2,2}}
     */
    private boolean failedCase2(int[][] matrix) {
      int lng = matrix.length;
      int total = lng * (lng + 1);
      for (int i = 0; i < lng; i++) {
        int sum = 0;
        // row
        for (int j = 0; j < lng; j++) {
          sum += matrix[i][j];
        }
        // column
        for (int j = 0; j < lng; j++) {
          sum += matrix[j][i];
        }
        if (sum != total) {
          return false;
        }
      }
      return true;
    }

    /**
     * {{1,2},{2,2}}
     */
    private boolean failedCase1(int[][] matrix) {
      int lng = matrix.length;
      for (int i = 0; i < lng; i++) {
        int xor = 0;
        // row
        for (int j = 0; j < lng; j++) {
          xor ^= matrix[i][j];
        }
        // column
        for (int j = 0; j < lng; j++) {
          xor ^= matrix[j][i];
        }
        if (xor != 0) {
          return false;
        }
      }
      return true;
    }
  }
//leetcode submit region end(Prohibit modification and deletion)


}