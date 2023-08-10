package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when234_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isPalindrome(NewList([]int{1, 2, 3, 2, 1})))
	assert.Equal(t, true, isPalindrome(NewList([]int{1, 2, 2, 1})))
	assert.Equal(t, false, isPalindrome(NewList([]int{1, 2})))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// length
	length := 0
	for ptr := head; ptr != nil; ptr = ptr.Next {
		length += 1
	}
	if length < 2 {
		return true
	}
	// 不用考虑长度为奇数或偶数
	ptr := head
	for i := 0; i < (length-1)/2; i++ {
		ptr = ptr.Next
	}
	// 反转后半段链表
	var pre *ListNode = nil
	for cur := ptr.Next; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 双指针
	ans := true
	left, right := head, pre
	for right != nil {
		if left.Val != right.Val {
			ans = false
			break
		}
		left, right = left.Next, right.Next
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
