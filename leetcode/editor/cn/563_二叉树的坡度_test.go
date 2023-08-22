package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when563_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{4, 2, 9, 3, 5, nil, 7})
	ans := findTilt(root)
	assert.Equal(t, 15, ans)
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
func findTilt(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)
		// 计算 node 的坡度
		if leftSum > rightSum {
			ans += leftSum - rightSum
		} else {
			ans += rightSum - leftSum
		}
		// 返回以 node 为根结点的 sum
		return leftSum + rightSum + node.Val
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
