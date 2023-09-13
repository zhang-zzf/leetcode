package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1022_thenSuccess(t *testing.T) {
	tree := decodeTreeNode([]any{1, 0, 1, 0, 1, 0, 1})
	assert.Equal(t, 22, sumRootToLeaf(tree))
}

func Test_givenFailedCase1_when1022_thenSuccess(t *testing.T) {
	tree := decodeTreeNode([]any{0, 1, 0, 0, nil, 0, 0, nil, nil, nil, 1, nil, nil, nil, 1})
	assert.Equal(t, 5, sumRootToLeaf(tree))
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
func sumRootToLeaf(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, v int) {
		if n == nil {
			return
		}
		v = v<<1 | n.Val
		if n.Left == nil && n.Right == nil {
			ans += v
		}
		dfs(n.Left, v)
		dfs(n.Right, v)
	}
	dfs(root, 0)
	return ans

}

func sumRootToLeafError(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, v int) {
		if n == nil {
			// 算法是错误的， n == nil 不代表是叶子节点
			ans += v
			return
		}
		v = v<<1 | n.Val
		dfs(n.Left, v)
		dfs(n.Right, v)
	}
	dfs(root, 0)
	return ans / 2
}

//leetcode submit region end(Prohibit modification and deletion)
