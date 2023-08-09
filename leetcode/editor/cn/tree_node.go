package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func encodeTreeNode(root *TreeNode) []any {
	var ans []any
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			ans = append(ans, nil)
		} else {
			ans = append(ans, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	// remove all the tail nil
	for i := len(ans) - 1; i >= 0 && ans[i] == nil; i-- {
		ans = ans[:i]
	}
	return ans
}

func decodeTreeNode(nodes []any) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	root := &TreeNode{Val: nodes[0].(int)}
	nodes = nodes[1:]
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		if r == nil {
			continue
		}
		leftFilled := false
		for i := 0; i < 2 && len(nodes) > 0; i++ {
			v := nodes[0]
			nodes = nodes[1:]
			// 三目 不存在
			var nn *TreeNode
			if v != nil {
				nn = &TreeNode{Val: v.(int)}
			}
			if !leftFilled {
				r.Left = nn
				leftFilled = true
			} else {
				r.Right = nn
			}
			queue = append(queue, nn)
		}
	}
	return root
}
