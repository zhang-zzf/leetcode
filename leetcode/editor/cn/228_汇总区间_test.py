import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ranges = self.solution.summaryRanges([0, 1, 2, 4, 5, 7])
        self.assertEqual(["0->2", "4->5", "7"], ranges)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        ans = []
        left, right = 0, 0
        while right < len(nums):
            if right == len(nums) - 1 or nums[right + 1] - nums[right] > 1:
                # [low,right]
                if right == left:
                    ans.append(str(nums[left]))
                else:
                    ans.append(str(nums[left]) + '->' + str(nums[right]))
                # reset the new left ptr
                left = right + 1
            right += 1

        return ans
        pass


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
