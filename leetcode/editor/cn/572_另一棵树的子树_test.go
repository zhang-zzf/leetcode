package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when572_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{3, 4, 5, 1, 2})
	subRoot := decodeTreeNode([]any{4, 1, 2})
	ans := isSubtree(root, subRoot)
	assert.Equal(t, true, ans)
}

func Test_givenFailedCase1_when572_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{3, 4, 5, 1, 2, nil, nil, nil, nil, 0})
	subRoot := decodeTreeNode([]any{4, 1, 2})
	ans := isSubtree(root, subRoot)
	assert.Equal(t, false, ans)
}

func Test_givenFailedCase2_when572_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{1, 1})
	subRoot := decodeTreeNode([]any{1})
	ans := isSubtree(root, subRoot)
	assert.Equal(t, true, ans)
}

func Test_givenFailedCase3_when572_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{3, 4, 5, 1, nil, 2})
	subRoot := decodeTreeNode([]any{3, 1, 2})
	ans := isSubtree(root, subRoot)
	assert.Equal(t, false, ans)
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfs572(root, subRoot) {
		return true
	}
	if isSubtree(root.Left, subRoot) {
		return true
	}
	if isSubtree(root.Right, subRoot) {
		return true
	}
	return false
}

func dfs572(n1 *TreeNode, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n2 == nil || n1 == nil {
		return false
	}
	// 都不为空
	if n1.Val == n2.Val {
		if !dfs572(n1.Left, n2.Left) {
			return false
		}
		if !dfs572(n1.Right, n2.Right) {
			return false
		}
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
