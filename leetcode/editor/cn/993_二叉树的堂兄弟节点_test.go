package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when993_thenSuccess(t *testing.T) {
	tree := decodeTreeNode([]any{1, 2, 3, nil, 4, nil, 5})
	assert.Equal(t, true, isCousins(tree, 4, 5))
	tree1 := decodeTreeNode([]any{1, 2, 3, 4})
	assert.Equal(t, false, isCousins(tree1, 4, 3))
}

func Test_givenFailedCase1_when993_thenSuccess(t *testing.T) {
	tree := decodeTreeNode([]any{1, 2, 3, nil, 4})
	assert.Equal(t, false, isCousins(tree, 2, 3))
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
// bfs
func isCousins1(root *TreeNode, x int, y int) bool {
	var level [][]*TreeNode
	level = append(level, []*TreeNode{root, nil})
	for len(level) > 0 {
		var xn, yn []*TreeNode
		// TODO 先计算 size
		size := len(level)
		// 放入 for 循环计算 size 的算法是错误的
		// for i := 0; i < len(level); i++ {
		for i := 0; i < size; i++ {
			first := level[0]
			level = level[1:]
			node := first[0]
			if node.Left != nil {
				level = append(level, []*TreeNode{node.Left, node})
			}
			if node.Right != nil {
				level = append(level, []*TreeNode{node.Right, node})
			}
			if node.Val == x {
				xn = first
			}
			if node.Val == y {
				yn = first
			}
		}
		if xn != nil && yn != nil {
			return xn[1] != yn[1]
		} else if xn != nil || yn != nil {
			return false
		}
	}
	return false
}

// dfs
func isCousins(root *TreeNode, x int, y int) bool {
	xk, xp, xf, yk, yp, yf := 0, root, false, 0, root, false
	var dfs func(*TreeNode, *TreeNode, int)
	dfs = func(n *TreeNode, p *TreeNode, depth int) {
		if n == nil {
			return
		}
		dfs(n.Left, n, depth+1)
		if n.Val == x {
			xk, xp, xf = depth, p, true
		}
		if n.Val == y {
			yk, yp, yf = depth, p, true
		}
		if xf && yf {
			return
		}
		dfs(n.Right, n, depth+1)
	}
	dfs(root, nil, 0)
	return xk == yk && xp != yp
}

//leetcode submit region end(Prohibit modification and deletion)
