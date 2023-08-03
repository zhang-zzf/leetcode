import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def reverseBits(self, n: int) -> int:
        # 16 -> 8 -> 4 -> 2 -> 1
        n = n >> 16 | n << 16
        n = (n & 0xff00ff00) >> 8 | (n & 0x00ff00ff) << 8
        n = (n & 0xf0f0f0f0) >> 4 | (n & 0x0f0f0f0f) << 4
        n = (n & 0xcccccccc) >> 2 | (n & 0x33333333) << 2
        n = (n & 0xaaaaaaaa) >> 1 | (n & 0x55555555) << 1
        return n
        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
