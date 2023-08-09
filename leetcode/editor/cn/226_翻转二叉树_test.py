import unittest
from typing import Optional

from model.tree_node import TreeNode


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if root:
            left = self.invertTree(root.left)
            right = self.invertTree(root.right)
            root.left, root.right = right, left
        return root


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
