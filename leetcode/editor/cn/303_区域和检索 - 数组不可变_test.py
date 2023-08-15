import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = NumArray()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


# leetcode submit region begin(Prohibit modification and deletion)
class NumArray:

    def __init__(self, nums: List[int]):
        # 前缀和
        lng = len(nums)
        pre_sum = [0] * (lng + 1)
        for x in range(1, lng + 1):
            pre_sum[x] = pre_sum[x - 1] + nums[x - 1]
        self.pre_sum = pre_sum

    def sumRange(self, left: int, right: int) -> int:
        return self.pre_sum[right + 1] - self.pre_sum[left]


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)
# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
