import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual(2, self.solution.addDigits(38))
        self.assertEqual(0, self.solution.addDigits(0))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def addDigits(self, num: int) -> int:
        ans = 0
        while num != 0:
            ans += num % 10
            num //= 10
        if ans < 10:
            return ans
        else:
            return self.addDigits(ans)


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
