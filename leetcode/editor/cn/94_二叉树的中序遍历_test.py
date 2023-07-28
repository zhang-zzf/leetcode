import unittest
from typing import Optional, List

from model.tree_node import TreeNode


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        # TODO List[int] 中可以接受 None
        root = TreeNode.decode([1, None, 2, 3])
        ans = self.solution.inorderTraversal(root)
        self.assertEqual([1, 3, 2], ans)


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        ans: list[any] = []
        while root:
            # left -> Node -> right
            if root.left:
                predecessor = root.left
                while predecessor.right and predecessor.right != root:
                    predecessor = predecessor.right
                if predecessor.right is None:
                    # iterate left
                    predecessor.right = root
                    root = root.left
                    continue
                else:
                    # reset tree
                    predecessor.right = None
            # Node
            ans.append(root.val)
            root = root.right
        return ans
        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
