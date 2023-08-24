package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when671_thenSuccess(t *testing.T) {
	assert.Equal(t, -1, findSecondMinimumValue(decodeTreeNode([]any{2, 2, 2})))
	assert.Equal(t, 5, findSecondMinimumValue(decodeTreeNode([]any{2, 2, 5})))
	assert.Equal(t, 4, findSecondMinimumValue(decodeTreeNode([]any{2, 2, 5, 4, 2})))
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
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	leftMin, rightMin := -1, -1
	if root.Val != root.Left.Val {
		leftMin = root.Left.Val
	} else {
		leftMin = findSecondMinimumValue(root.Left)
	}
	if root.Val != root.Right.Val {
		rightMin = root.Right.Val
	} else {
		rightMin = findSecondMinimumValue(root.Right)
	}
	if leftMin == -1 {
		return rightMin
	}
	if rightMin == -1 {
		return leftMin
	}
	return int(math.Min(float64(leftMin), float64(rightMin)))
}

//leetcode submit region end(Prohibit modification and deletion)
