import model.TreeNode;
import org.junit.jupiter.api.Test;


class EvaluateBooleanBinaryTreeTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {

    }

    // leetcode submit region begin(Prohibit modification and deletion)

    /**
     * Definition for a binary tree node. public class TreeNode { int val; TreeNode left; TreeNode right; TreeNode() {}
     * TreeNode(int val) { this.val = val; } TreeNode(int val, TreeNode left, TreeNode right) { this.val = val;
     * this.left = left; this.right = right; } }
     */
    class Solution {
        public boolean evaluateTree(TreeNode root) {
            if (root.left == null) {
                return root.val == 1;
            }
            return switch (root.val) {
                case 2 -> evaluateTree(root.left) || evaluateTree(root.right);
                case 3 -> evaluateTree(root.left) && evaluateTree(root.right);
                default -> false;
            };
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}