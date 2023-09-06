package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when897_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9})
	ans := []any{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9}
	bst := increasingBST(root)
	assert.Equal(t, ans, encodeTreeNode(bst))
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
// TODO 参考答案，中序遍历
func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	ptr := dummy
	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		// deal node start
		ptr.Right, ptr = n, n
		n.Left = nil
		// deal node end
		inorder(n.Right)
	}
	inorder(root)
	return dummy.Right
}

// 后序遍历
func increasingBST11(root *TreeNode) *TreeNode {
	var postOrder func(*TreeNode) (*TreeNode, *TreeNode)
	postOrder = func(n *TreeNode) (*TreeNode, *TreeNode) {
		if n == nil {
			return nil, nil
		}
		f, t := n, n
		if n.Left != nil {
			lf, lt := postOrder(n.Left)
			f = lf
			lt.Right = n
			n.Left = nil
		}
		if n.Right != nil {
			rl, rt := postOrder(n.Right)
			t = rt
			n.Right = rl
		}
		return f, t
	}
	head, _ := postOrder(root)
	return head
}

// 先序遍历实现
func increasingBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// pre-order-traversal 先序遍历
	l, r := root.Left, root.Right
	if root.Left != nil {
		n := root.Left
		for n != nil && n.Right != nil {
			n = n.Right
		}
		n.Right = root
		root.Left = nil
	}
	if root.Right != nil {
		n := root.Right
		for n != nil && n.Left != nil {
			n = n.Left
		}
		root.Right = n
	}
	// 左子树中查找 head
	head := increasingBST(l)
	increasingBST(r)
	if head == nil {
		return root
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
