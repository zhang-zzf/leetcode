package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when965_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{2, 2, 2, 5, 2})
	assert.Equal(t, false, isUnivalTree(root))
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
func isUnivalTree(root *TreeNode) bool {
	var dfs func(*TreeNode) bool
	dfs = func(n *TreeNode) bool {
		if n == nil {
			return true
		}
		// 先序遍历
		if n.Val != root.Val {
			return false
		}
		return dfs(n.Left) && dfs(n.Right)
	}
	return dfs(root)
}

//leetcode submit region end(Prohibit modification and deletion)
