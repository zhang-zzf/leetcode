package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when404_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{3, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, 24, sumOfLeftLeaves(root))
	root = decodeTreeNode([]any{1})
	assert.Equal(t, 0, sumOfLeftLeaves(root))
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
func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesRec(root, false)
}

func sumOfLeftLeavesRec(root *TreeNode, left bool) int {
	if root == nil {
		return 0
	}
	if left && root.Left == nil && root.Right == nil {
		return root.Val
	}
	leftSum := sumOfLeftLeavesRec(root.Left, true)
	rightSum := sumOfLeftLeavesRec(root.Right, false)
	return leftSum + rightSum
}

//leetcode submit region end(Prohibit modification and deletion)
