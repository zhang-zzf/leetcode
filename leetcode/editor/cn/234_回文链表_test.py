import unittest
from typing import Optional

from model.list_node import ListNode


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        palindrome = ListNode.new_list([1, 2, 3, 2, 1])
        self.assertTrue(self.solution.isPalindrome(palindrome))
        #
        palindrome = ListNode.new_list([1, 2, 2, 1])
        self.assertTrue(self.solution.isPalindrome(palindrome))
        #
        not_palindrome = ListNode.new_list([1, 2])
        self.assertFalse(self.solution.isPalindrome(not_palindrome))

# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        if head is None or head.next is None:
            return True
        stack = []
        slow, fast = head, head.next
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
        # slow is the first half
        ptr = head
        while ptr != slow.next:
            stack.append(ptr.val)
            ptr = ptr.next
        # 奇数长度链表
        if fast is None:
            stack.pop()
        while ptr is not None and ptr.val == stack.pop():
            ptr = ptr.next
        return ptr is None


# leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    unittest.main()
