//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树
// 👍 1095 👎 0


import model.TreeNode;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.assertj.core.api.BDDAssertions.then;


class BinaryTreePreorderTraversalTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        TreeNode root = TreeNode.decode("[1,null,2,3]");
        List<Integer> ans = solution.preorderTraversal(root);
        then(ans).containsExactly(1, 2, 3);
    }

    //leetcode submit region begin(Prohibit modification and deletion)

    /**
     * Definition for a binary tree node.
     * public class TreeNode {
     * int val;
     * TreeNode left;
     * TreeNode right;
     * TreeNode() {}
     * TreeNode(int val) { this.val = val; }
     * TreeNode(int val, TreeNode left, TreeNode right) {
     * this.val = val;
     * this.left = left;
     * this.right = right;
     * }
     * }
     */
    class Solution {

        public List<Integer> preorderTraversal(TreeNode root) {
            List<Integer> ans = new ArrayList<>();
            while (root != null) {
                // preorder traversal
                // Node -> left -> right
                // Node
                ans.add(root.val);
                // left
                if (root.left != null) {
                    var predecessor = root.left;
                    while (predecessor.right != null && predecessor.right != root) {
                        predecessor = predecessor.right;
                    }
                    if (predecessor.right == null) {
                        predecessor.right = root;
                        // move pointer
                        root = root.left;
                        continue;
                    } else {
                        predecessor.right = null;
                        ans.remove(ans.size() - 1);
                    }
                }
                // inorder traversal
                //
                // right
                root = root.right;
            }
            return ans;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)


}