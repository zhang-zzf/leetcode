package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when501_thenSuccess(t *testing.T) {
	assert.Equal(t, []int{2}, findMode(decodeTreeNode([]any{1, nil, 2, 2})))
}

/**
1. 边界
2.

*/
func Test_givenFailedCase1_when501_thenSuccess(t *testing.T) {
	/**
	> 2023/08/21 12:49:32
	Wrong Answer
	input:
	[0]
	Output:
	[0,0]
	Expected:
	[0]
	stdout:


	> 2023/08/21 12:50:01
	Code submitted. Please wait...

	> 2023/08/21 12:50:03
	Wrong Answer
	input:
	[1,0,1,0,0,1,1,0]
	Output:
	[1]
	Expected:
	[0,1]
	stdout:
	*/
	assert.Equal(t, []int{2}, findMode(decodeTreeNode([]any{1, nil, 2, 2})))
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
func findMode(root *TreeNode) []int {
	var ans []int
	var pre *TreeNode
	cnt, maxCount := 1, 1
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		//
		if pre != nil && n.Val == pre.Val {
			cnt += 1
		} else {
			cnt = 1
		}
		if cnt == maxCount {
			ans = append(ans, n.Val)
		} else if cnt > maxCount {
			ans = []int{n.Val}
			maxCount = cnt
		}
		pre = n
		//
		dfs(n.Right)
	}
	dfs(root)
	return ans
}

func findMode1(root *TreeNode) []int {
	var ans []int
	var preNode *TreeNode
	cnt, maxCnt := 1, 1
	for root != nil {
		if root.Left != nil {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right != root {
				predecessor.Right = root
				root = root.Left
				continue
			} else {
				predecessor.Right = nil
			}
		}
		// inorder-traversal
		val := root.Val
		if preNode != nil && val == preNode.Val {
			cnt += 1
		} else {
			cnt = 1
		}
		if cnt == maxCnt {
			ans = append(ans, val)
		} else if cnt > maxCnt {
			ans = []int{val}
			maxCnt = cnt
		}
		preNode = root
		//
		root = root.Right
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
