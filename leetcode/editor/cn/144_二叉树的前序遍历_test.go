package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when144_thenSuccess(t *testing.T) {
	tree := decodeTreeNode([]any{1, nil, 2, 3})
	ans := preorderTraversal(tree)
	assert.Equal(t, []int{1, 2, 3}, ans, "shouldEqual")
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
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	for root != nil {
		// Node -> left -> right
		ans = append(ans, root.Val)
		// left
		if root.Left != nil {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
				continue
			} else {
				predecessor.Right = nil
				ans = ans[:len(ans)-1]
			}
		}
		// right
		root = root.Right
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
