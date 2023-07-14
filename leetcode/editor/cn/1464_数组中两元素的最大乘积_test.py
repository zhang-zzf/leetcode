import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        m1, m2 = nums[0], nums[1]
        if m1 < m2:
            m1, m2 = m2, m1
        for n in nums[2:]:
            if n > m1:
                m1, m2 = n, m1
            elif n > m2:
                m2 = n
        return (m1 - 1) * (m2 - 1)


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
