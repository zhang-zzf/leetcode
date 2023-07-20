package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when160_thenSuccess(t *testing.T) {
	n3 := &ListNode{3, nil}
	h1 := &ListNode{1, &ListNode{2, n3}}
	h2 := &ListNode{4, n3}
	assert.Equal(t, n3, getIntersectionNode(h1, h2), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	// nil == nil is true
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}

//leetcode submit region end(Prohibit modification and deletion)
