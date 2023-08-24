package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when617_thenSuccess(t *testing.T) {
	r1 := decodeTreeNode([]any{1, 3, 2, 5})
	r2 := decodeTreeNode([]any{2, 1, 3, nil, 4, nil, 7})
	ans := encodeTreeNode(mergeTrees(r1, r2))
	assert.Equal(t, []any{3, 4, 5, 5, 4, nil, 7}, ans)
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	v1, v2 := 0, 0
	var l1, l2, r1, r2 *TreeNode
	if root1 != nil {
		v1, l1, r1 = root1.Val, root1.Left, root1.Right
	}
	if root2 != nil {
		v2, l2, r2 = root2.Val, root2.Left, root2.Right
	}
	root := &TreeNode{Val: v1 + v2}
	root.Left = mergeTrees(l1, l2)
	root.Right = mergeTrees(r1, r2)
	return root
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{root1.Val + root2.Val,
		mergeTrees1(root1.Left, root2.Left),
		mergeTrees1(root1.Right, root2.Right),
	}
}

//leetcode submit region end(Prohibit modification and deletion)
