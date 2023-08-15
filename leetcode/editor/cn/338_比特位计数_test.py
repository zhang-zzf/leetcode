import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        bits = self.solution.countBits(5)
        self.assertEqual([0, 1, 1, 2, 1, 2], bits)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def countBits(self, n: int) -> List[int]:
        ans = [0] * (n + 1)
        for x in range(0, n + 1):
            if x % 2 == 0:
                ans[x] = ans[x // 2]
            else:
                ans[x] = ans[x - 1] + 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
