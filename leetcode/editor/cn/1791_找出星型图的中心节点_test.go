package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1791_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findCenter(edges [][]int) int {
	e1, e2 := edges[0], edges[1]
	if e1[0] == e2[0] || e1[0] == e2[1] {
		return e1[0]
	}
	return e1[1]
}

func findCenter1(edges [][]int) int {
	n := len(edges)
	cnt := make([]int, n+2)
	for _, edge := range edges {
		cnt[edge[0]] += 1
		cnt[edge[1]] += 1
	}
	for node, c := range cnt {
		if c == n {
			return node
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
