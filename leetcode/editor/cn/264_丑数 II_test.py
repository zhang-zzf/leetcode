import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertEqual(12, self.solution.nthUglyNumber(10))


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def nthUglyNumber(self, n: int) -> int:
        # TODO 创建一定长度的数组
        ugly = [1] * (n + 1)
        p2, p3, p5 = 1, 1, 1
        for x in range(2, n + 1):
            n2, n3, n5 = ugly[p2] * 2, ugly[p3] * 3, ugly[p5] * 5
            next_ugly = min(n2, n3, n5)
            ugly[x] = next_ugly
            # watch out
            if next_ugly == n2:
                p2 += 1
            if next_ugly == n3:
                p3 += 1
            if next_ugly == n5:
                p5 += 1
        return ugly[n]
