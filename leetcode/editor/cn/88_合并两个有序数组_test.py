import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        nums: List[int] = [1, 2, 3, 0, 0, 0]
        self.solution.merge(nums, 3, [2, 5, 6], 3)
        self.assertListEqual(nums, [1, 2, 2, 3, 5, 6])


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:

    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        for idx in reversed(range(len(nums1))):
            if n - 1 < 0:
                break
            if m - 1 >= 0 and nums1[m - 1] >= nums2[n - 1]:
                nums1[idx] = nums1[m - 1]
                m -= 1
            else:
                nums1[idx] = nums2[n - 1]
                n -= 1


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
