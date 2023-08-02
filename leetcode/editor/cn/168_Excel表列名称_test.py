import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual("AZ", self.solution.convertToTitle(52))
        self.assertEqual("A", self.solution.convertToTitle(1))
        self.assertEqual("FXSHRXW", self.solution.convertToTitle(2147483647))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        BASE = 26
        ans = ''
        while columnNumber:
            columnNumber -= 1
            ans = chr(ord('A') + columnNumber % BASE) + ans
            columnNumber //= BASE
        return ans
        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
