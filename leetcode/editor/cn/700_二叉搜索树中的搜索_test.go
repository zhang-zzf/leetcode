package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when700_thenSuccess(t *testing.T) {
	ans := encodeTreeNode(searchBST(decodeTreeNode([]any{4, 2, 7, 1, 3}), 2))
	assert.Equal(t, []any{2, 1, 3}, ans)
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
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val == val:
		return root
	case root.Val < val:
		return searchBST(root.Right, val)
	case root.Val > val:
		return searchBST(root.Left, val)
	default:
		return nil
	}
}

//leetcode submit region end(Prohibit modification and deletion)
