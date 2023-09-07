package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when938_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{10, 5, 15, 3, 7, 13, 18, 1, nil, 6})
	ans := rangeSumBST(root, 6, 10)
	assert.Equal(t, 23, ans)
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

// TODO 参考答案
func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Val >= low {
			inorder(n.Left)
		}
		// inorder start
		if n.Val >= low && n.Val <= high {
			ans += n.Val
		}
		// inorder end
		if n.Val <= high {
			inorder(n.Right)
		}
	}
	inorder(root)
	return ans
}

func rangeSumBST1(root *TreeNode, low int, high int) int {
	ans := 0
	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		// inorder start
		if n.Val >= low && n.Val <= high {
			ans += n.Val
		}
		// inorder end
		inorder(n.Right)
	}
	inorder(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
