import unittest
from typing import Optional

from model.list_node import ListNode


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        n3 = ListNode(3)
        h1 = ListNode(1, ListNode(2, n3))
        h2 = ListNode(4, n3)
        node = self.solution.getIntersectionNode(h1, h2)
        self.assertEqual(n3, node)


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        pa, pb = headA, headB
        while pa != pb:
            pa = pa.next if pa else headB
            pb = pb.next if pb else headA
        return pa


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
