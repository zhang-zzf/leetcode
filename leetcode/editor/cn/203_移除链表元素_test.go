package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when203_thenSuccess(t *testing.T) {
	list := NewListFromSlice([]int{7, 7, 7})
	assert.Nil(t, removeElements(list, 7))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	for ptr := dummy; ptr.Next != nil; {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
