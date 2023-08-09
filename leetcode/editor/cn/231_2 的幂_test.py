import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(self.solution.isPowerOfTwo(1))
        self.assertTrue(self.solution.isPowerOfTwo(16))
        self.assertFalse(self.solution.isPowerOfTwo(15))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        pivot = 1
        while pivot < n:
            pivot = pivot << 1
        return pivot == n


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
