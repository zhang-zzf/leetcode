package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
		for i := 0; i < Min(2, len(nodes)); i++ {
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

func Min[T int](args ...T) T {
	min := args[0]
	for _, x := range args {
		if x < min {
			min = x
		}
	}
	return min
}

func Max[T int](args ...T) T {
	max := args[0]
	for _, x := range args {
		if x > max {
			max = x
		}
	}
	return max
}
