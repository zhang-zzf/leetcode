import unittest
from typing import Optional, List

from model.tree_node import TreeNode


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        root = TreeNode.decode([1, 2, 3, None, 5])
        paths = self.solution.binaryTreePaths(root)
        self.assertListEqual(["1->2->5", "1->3"], paths)


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        ans = []
        if root is None:
            return ans
        if root.left:
            for item in self.binaryTreePaths(root.left):
                ans.append(str(root.val) + "->" + item)
        if root.right:
            for item in self.binaryTreePaths(root.right):
                ans.append(str(root.val) + "->" + item)
        if root.left is None and root.right is None:
            ans.append(str(root.val))
        return ans
        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
