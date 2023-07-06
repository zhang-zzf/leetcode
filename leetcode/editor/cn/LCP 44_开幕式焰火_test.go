package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_whenLCP44_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func numColor(root *TreeNode) int {
	colors := make(map[int]struct{})
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		colors[root.Val] = struct{}{}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return len(colors)
}

//leetcode submit region end(Prohibit modification and deletion)
