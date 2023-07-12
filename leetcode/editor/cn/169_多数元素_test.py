import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ans = self.solution.majorityElement([2, 1, 2, 3, 2, 3, 2])
        self.assertEqual(2, ans)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        ans, cnt = 0, 0
        for n in nums:
            if cnt == 0:
                ans = n
                cnt += 1
            elif n == ans:
                cnt += 1
            else:
                cnt -= 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
