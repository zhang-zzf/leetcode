import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual(self.solution.hammingWeight(3), 2)
        self.assertEqual(self.solution.hammingWeight(-1), 32)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def hammingWeight(self, n: int) -> int:
        # 无法统计负数
        # return f'{n:b}'.count('1')
        ans: int = 0
        for x in range(32):
            if n == 0:
                break
            n &= n - 1
            ans += 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
