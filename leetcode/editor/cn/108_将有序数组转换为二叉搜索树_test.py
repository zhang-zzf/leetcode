import unittest
from typing import List, Optional

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
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        if len(nums) == 0:
            return None

        # TODO len(nums)/2 result is float
        # pivot: int = int(len(nums) / 2)
        pivot: int = len(nums) // 2
        return TreeNode(nums[pivot],
                        self.sortedArrayToBST(nums[:pivot]),
                        self.sortedArrayToBST(nums[pivot + 1:])
                        )

        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
