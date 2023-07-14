package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when145_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{1, nil, 2, 3, nil, 5})
	ans := postorderTraversal(root)
	assert.Equal(t, []int{5, 3, 2, 1}, ans, "shouldEqual")
}

func Test_givenNormal_whenDecodeTreeNode_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{1, nil, 2, 3, nil, 5})
	assert.Equal(t, 2, root.Right.Val, "shouldEqual")
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	// TODO nil... 可以被展开
	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
