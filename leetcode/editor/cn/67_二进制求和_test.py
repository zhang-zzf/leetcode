import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ans = self.solution.addBinary("11", "111")
        self.assertEqual("1010", ans)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        ans: str = ""
        promotion: int = 0
        ia, ib = len(a) - 1, len(b) - 1
        for _ in range(max(len(a), len(b))):
            na, nb = 0, 0
            if ia >= 0:
                na = int(a[ia])
            if ib >= 0:
                nb = int(b[ib])
            sum_val: int = na + nb + promotion
            ans = str(sum_val % 2) + ans
            promotion = 1 if sum_val > 1 else 0
            ia, ib = ia - 1, ib - 1
        if promotion:
            ans = "1" + ans
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
