import unittest
from typing import Optional, List

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
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        ans: list[int] = []
        while root:
            # Node -> left -> right
            ans.append(root.val)
            if root.left:
                predecessor = root.left
                while predecessor.right and predecessor.right != root:
                    predecessor = predecessor.right
                if predecessor.right is None:
                    predecessor.right = root
                    root = root.left
                    continue
                else:
                    # reset tree
                    predecessor.right = None
                    ans.pop()
            root = root.right
        return ans
        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
