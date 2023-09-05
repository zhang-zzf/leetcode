package leetcode

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when872_thenSuccess(t *testing.T) {
	root1 := decodeTreeNode([]any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4})
	root2 := decodeTreeNode([]any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8})
	assert.Equal(t, true, leafSimilar(root1, root2))
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	findLeaf := func(root *TreeNode) string {
		var ans []int
		var dfs func(node *TreeNode)
		dfs = func(node *TreeNode) {
			if node == nil {
				return
			}
			// inorder-traversal
			dfs(node.Left)
			if node.Left == nil && node.Right == nil {
				ans = append(ans, node.Val)
			}
			dfs(node.Right)
		}
		dfs(root)
		jsonBytes, _ := json.Marshal(ans)
		return string(jsonBytes)
	}
	return findLeaf(root1) == findLeaf(root2)
}

//leetcode submit region end(Prohibit modification and deletion)
