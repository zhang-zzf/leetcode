package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromSlice(nums []int) *ListNode {
	dummy := &ListNode{Val: 0}
	ptr := dummy
	for _, num := range nums {
		ptr.Next = &ListNode{num, nil}
		ptr = ptr.Next
	}
	return dummy.Next
}
