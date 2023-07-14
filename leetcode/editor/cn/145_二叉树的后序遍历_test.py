import unittest
from typing import Optional, List

from model.tree_node import TreeNode


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        tree_node = TreeNode.decode([1, None, 2, 3, None, 5])
        ans = self.solution.postorderTraversal(tree_node)
        self.assertEqual([5, 3, 2, 1], ans)

    def test_givenDecode_when_thenSuccess(self):
        tree_node = TreeNode.decode([1, None, 2, 3, None, 5])
        self.assertEqual(1, tree_node.val)


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        # Optional 用法
        # root can be None or a reference
        if not root:
            return []
        left = self.postorderTraversal(root.left)
        right = self.postorderTraversal(root.right)
        # merge list
        return left + right + [root.val]


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
