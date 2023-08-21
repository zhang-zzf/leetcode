package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when530_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{4, 2, 6, 1, 3})
	ans := getMinimumDifference(root)
	assert.Equal(t, 1, ans)
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
func getMinimumDifference(root *TreeNode) int {
	// 递归 dfs
	ans, pre := math.MaxInt, -1
	// TODO 先声明再递归
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			diff := node.Val - pre
			if diff < ans {
				ans = diff
			}
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func getMinimumDifference1(root *TreeNode) int {
	// Morris 遍历
	ans := math.MaxInt
	var preNode *TreeNode
	for root != nil {
		if root.Left != nil {
			p := root.Left
			for p.Right != nil && p.Right != root {
				p = p.Right
			}
			if p.Right != root {
				p.Right = root
				root = root.Left
				continue
			} else {
				p.Right = nil
			}
		}
		val := root.Val
		if preNode != nil {
			subtract := val - preNode.Val
			if subtract < ans {
				ans = subtract
			}
		}
		preNode = root
		root = root.Right
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
