import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(self.solution.canWinNim(1), "ShouldBeTrue")
        self.assertFalse(self.solution.canWinNim(4), "ShouldBeFalse")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def canWinNim(self, n: int) -> bool:
        return (n % 4) != 0


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
