import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        missing = self.solution.missingNumber([3, 0, 1])
        self.assertEqual(2, missing)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        ans = 0
        for x in nums:
            ans ^= x
        for x in range(len(nums) + 1):
            ans ^= x
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
