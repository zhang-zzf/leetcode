package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when226_thenSuccess(t *testing.T) {
	tree := decodeTreeNode([]any{4, 2, 7, 1, 3, 6, 9})
	invertedTree := invertTree(tree)
	assert.Equal(t, []any{4, 7, 2, 9, 6, 3, 1}, encodeTreeNode(invertedTree))
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
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		// switch left and right
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
