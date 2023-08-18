import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        ans = []
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        cnt_map = {}
        for x in nums1:
            cnt_map[x] = 1 if x not in cnt_map else cnt_map[x] + 1
        for x in nums2:
            if x in cnt_map and cnt_map[x] > 0:
                ans.append(x)
                cnt_map[x] -= 1
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
