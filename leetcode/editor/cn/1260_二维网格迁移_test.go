package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1260_thenSuccess(t *testing.T) {
	params := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ans := shiftGrid(params, 9)
	assert.Equal(t, params, ans)
}

func Test_givenFailedCase1_when1260_thenSuccess(t *testing.T) {
	params := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ans := shiftGrid(params, 1)
	assert.Equal(t, [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}, ans)
}

func Test_givenFailedCase2_when1260_thenSuccess(t *testing.T) {
	params := [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}
	ans := shiftGrid(params, 23)
	assert.Equal(t, [][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	totalNum := m * n
	idx := k % totalNum
	ans := make([][]int, m)
	for r := range ans {
		ans[r] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// watch out for idx overflow
			// watch out for i*n not i*m
			ti := (i*n + j + idx) % totalNum
			// watch out for ti/n not ti/m
			ans[(ti / n)][(ti % n)] = grid[i][j]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
