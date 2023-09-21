package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1351_thenSuccess(t *testing.T) {
	params := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	ans := countNegatives(params)
	assert.Equal(t, 8, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for r, c := 0, n-1; r < m; r++ {
		for c >= 0 && grid[r][c] < 0 {
			c -= 1
		}
		ans += n - c - 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
