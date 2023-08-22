package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when543_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{1, 2, 3, 4, 5})
	ans := diameterOfBinaryTree(root)
	assert.Equal(t, 3, ans)
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
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := 0, 0
		if node.Left != nil {
			left = 1 + dfs(node.Left)
		}
		if node.Right != nil {
			right = 1 + dfs(node.Right)
		}
		diameter := left + right
		if diameter > ans {
			ans = diameter
		}
		// return border
		border := left
		if right > border {
			border = right
		}
		return border
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
