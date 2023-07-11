import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ss = self.solution.minSubsequence([4, 3, 10, 9, 8])
        self.assertListEqual([10, 9], ss)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def minSubsequence(self, nums: List[int]) -> List[int]:
        nums.sort(reverse=True)
        right, left = sum(nums) / 2, 0
        for i in range(len(nums)):
            left += nums[i]
            if left > right:
                return nums[:i + 1]


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
