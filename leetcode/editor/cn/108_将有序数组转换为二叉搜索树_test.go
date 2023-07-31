package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when108_thenSuccess(t *testing.T) {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	nodes := encodeTreeNode(root)
	assert.Equal(t, []any{0, -3, 9, -10, nil, 5}, nodes)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	pivot := len(nums) / 2
	left := sortedArrayToBST(nums[0:pivot])
	right := sortedArrayToBST(nums[pivot+1:])
	return &TreeNode{Val: nums[pivot], Left: left, Right: right}
}

//leetcode submit region end(Prohibit modification and deletion)
