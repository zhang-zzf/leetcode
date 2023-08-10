import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isUgly(self, n: int) -> bool:
        factors = [2, 3, 5]
        while n > 1:
            ugly = False
            for f in factors:
                if n % f == 0:
                    n /= f
                    ugly = True
            if not ugly:
                break
        return n == 1


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
