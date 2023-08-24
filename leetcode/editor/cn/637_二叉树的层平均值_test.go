package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when637_thenSuccess(t *testing.T) {
	root := decodeTreeNode([]any{3, 9, 20, 15, 7})
	ans := averageOfLevels(root)
	assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, ans)
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
func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	var queue []*TreeNode
	// 单队列遍历 level
	queue = append(queue, root)
	for len(queue) > 0 {
		size, sum := len(queue), 0
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			sum += n.Val
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(size))
	}
	return ans
}

func averageOfLevels1(root *TreeNode) []float64 {
	var ans []float64
	if root == nil {
		return ans
	}
	// TODO 使用2个队列表示 level -> nextLevel
	var cl, nl []*TreeNode
	cl = append(cl, root)
	cs, cc := 0, 0
	for len(cl) != 0 {
		n := cl[0]
		cl = cl[1:]
		if n != nil {
			nl = append(nl, n.Left, n.Right)
			cs, cc = cs+n.Val, cc+1
		}
		// 当前 level 已经遍历完成，且存在下一 level
		if len(cl) == 0 && len(nl) != 0 {
			ans = append(ans, float64(cs)/float64(cc))
			cl, nl = nl, cl
			cs, cc = 0, 0
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
