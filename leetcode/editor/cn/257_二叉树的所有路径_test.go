package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when257_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
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
func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	if root == nil {
		return ans
	}
	rootVal := strconv.Itoa(root.Val)
	if root.Left != nil {
		for _, item := range binaryTreePaths(root.Left) {
			ans = append(ans, rootVal+"->"+item)
		}
	}
	if root.Right != nil {
		for _, item := range binaryTreePaths(root.Right) {
			ans = append(ans, rootVal+"->"+item)
		}
	}
	if root.Left == nil && root.Right == nil {
		ans = append(ans, rootVal)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
