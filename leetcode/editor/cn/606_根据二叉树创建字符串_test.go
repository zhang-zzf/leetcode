package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when606_thenSuccess(t *testing.T) {
	assert.Equal(t, "1(2(4))(3)", tree2str(decodeTreeNode([]any{1, 2, 3, 4})))
	assert.Equal(t, "1(2()(4))(3)", tree2str(decodeTreeNode([]any{1, 2, 3, nil, 4})))
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
func tree2str(root *TreeNode) string {
	// TODO switch
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}
func tree2str1(root *TreeNode) string {
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "()"
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if right == "()" {
			if left == "()" {
				return fmt.Sprintf("(%d)", node.Val)
			} else {
				return fmt.Sprintf("(%d%s)", node.Val, left)
			}
		} else {
			return fmt.Sprintf("(%d%s%s)", node.Val, left, right)
		}
	}
	ans := dfs(root)
	return ans[1 : len(ans)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
