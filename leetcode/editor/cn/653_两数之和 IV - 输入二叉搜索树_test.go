package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when653_thenSuccess(t *testing.T) {
	assert.Equal(t, true, findTarget(decodeTreeNode([]any{5, 3, 6, 2, 4, nil, 7}), 9))
	assert.Equal(t, false, findTarget(decodeTreeNode([]any{5, 3, 6, 2, 4, nil, 7}), 28))
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
// dfs
func findTarget(root *TreeNode, k int) bool {
	existVal := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(n *TreeNode) bool {
		if n == nil {
			return false
		}
		if dfs(n.Left) {
			return true
		}
		// inorder-traversal
		if _, ok := existVal[k-n.Val]; ok {
			return true
		}
		existVal[n.Val] = struct{}{}
		return dfs(n.Right)
	}
	return dfs(root)
}

// Morris 遍历
func findTarget1(root *TreeNode, k int) bool {
	ans := false
	existVal := map[int]struct{}{}
	for root != nil {
		if root.Left != nil {
			p := root.Left
			for p.Right != nil && p.Right != root {
				p = p.Right
			}
			if p.Right != root {
				p.Right = root
				root = root.Left
				continue
			}
			p.Right = nil
		}
		// inorder-traversal
		if _, ok := existVal[k-root.Val]; ok {
			ans = true
			break
		}
		existVal[root.Val] = struct{}{}
		//
		root = root.Right
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
