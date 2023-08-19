package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when463_thenSuccess(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
	}
	assert.Equal(t, 16, islandPerimeter(grid))
}

//leetcode submit region begin(Prohibit modification and deletion)
func islandPerimeter(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	// row then column
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			ans += 4
			// top / right / down / left
			if i-1 >= 0 && grid[i-1][j] == 1 {
				ans -= 1
			}
			if j+1 < n && grid[i][j+1] == 1 {
				ans -= 1
			}
			if i+1 < m && grid[i+1][j] == 1 {
				ans -= 1
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				ans -= 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
