import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
def dfs():
    pass


class Solution:
    def numColor(self, root: TreeNode) -> int:
        colors: set[int] = set()

        def dfs(root: TreeNode, colors: set[int]):
            if not root:
                return
            colors.add(root.val)
            dfs(root.left, colors)
            dfs(root.right, colors)

        dfs(root, colors)
        return len(colors)

        # leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
