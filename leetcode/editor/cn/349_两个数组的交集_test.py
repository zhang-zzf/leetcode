import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:

    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        ans = []
        base_cnt = {}
        for x in nums1:
            base_cnt[x] = 1
        for x in nums2:
            if x in base_cnt:
                del base_cnt[x]
                ans.append(x)
        return ans

    def intersection1(self, nums1: List[int], nums2: List[int]) -> List[int]:
        base_cnt = {}
        for x in nums1:
            base_cnt[x] = 1
        for x in nums2:
            if x in base_cnt:
                base_cnt[x] += 1
        ans = []
        # TODO 遍历 map
        for k, v in base_cnt.items():
            if v > 1:
                ans.append(k)
        return ans


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
