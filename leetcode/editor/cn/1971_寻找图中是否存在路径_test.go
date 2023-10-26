package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1971_thenSuccess(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	assert.Equal(t, false, validPath(6, edges, 0, 5))
}

func Test_givenFailedCase1_when1971_thenSuccess(t *testing.T) {
	edges := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
	assert.Equal(t, true, validPath(10, edges, 5, 9))
}

// todo dfs / 并查集 算法
//leetcode submit region begin(Prohibit modification and deletion)
func validPath(n int, edges [][]int, source int, destination int) bool {
	if source > destination {
		source, destination = destination, source
	}
	srcReachable := make([]int, n)
	mapping := map[int]map[int]struct{}{}
	for _, e := range edges {
		if m, ok := mapping[e[0]]; !ok {
			mapping[e[0]] = map[int]struct{}{e[1]: {}}
		} else {
			m[e[1]] = struct{}{}
		}
		if m, ok := mapping[e[1]]; !ok {
			mapping[e[1]] = map[int]struct{}{e[0]: {}}
		} else {
			m[e[0]] = struct{}{}
		}
	}
	var queue = []int{source}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		srcReachable[node] = 1
		for r := range mapping[node] {
			if srcReachable[r] == 0 {
				queue = append(queue, r)
			}
		}
	}
	return srcReachable[destination] == 1
}

/**
单向图的算法是错误的
*/
func validPathError(n int, edges [][]int, source int, destination int) bool {
	// 建立 小 -> 大 的单项图
	if source > destination {
		source, destination = destination, source
	}
	srcReachable := make([]int, n)
	mapping := map[int]map[int]struct{}{}
	for _, e := range edges {
		e0, e1 := e[0], e[1]
		if e0 > e1 {
			e0, e1 = e1, e0
		}
		if m, ok := mapping[e0]; !ok {
			mapping[e0] = map[int]struct{}{e1: {}}
		} else {
			m[e1] = struct{}{}
		}
	}
	var queue = []int{source}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		srcReachable[node] = 1
		for r := range mapping[node] {
			if srcReachable[r] == 0 {
				queue = append(queue, r)
			}
		}
	}
	return srcReachable[destination] == 1
}

//leetcode submit region end(Prohibit modification and deletion)
