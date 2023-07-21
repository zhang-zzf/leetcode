import unittest
from typing import List


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        ans = self.solution.nextGreaterElement([1, 4, 100],
                                               [1, 10, 4, 5, 2, 100])
        self.assertEqual([10, 5, -1], ans)


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def nextGreaterElement(self,
                           nums1: List[int],
                           nums2: List[int]) -> List[int]:
        mapping: dict[int, int] = {}
        stack: list[int] = []
        # reversed()
        for x in reversed(nums2):
            while stack and stack[- 1] <= x:
                stack.pop()
            mapping[x] = stack[- 1] if stack else -1
            stack.append(x)
        return [mapping[x] for x in nums1]


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
