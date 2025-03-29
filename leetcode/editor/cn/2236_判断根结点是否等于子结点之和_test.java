import model.TreeNode;
import org.junit.jupiter.api.Test;


class RootEqualsSumOfChildrenTest {

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
        public boolean checkTree(TreeNode root) {
            return root.val == root.left.val + root.right.val;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}