package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when892_thenSuccess(t *testing.T) {
	params := [][]int{{1, 2}, {3, 4}}
	assert.Equal(t, 34, surfaceArea(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func surfaceArea(grid [][]int) int {
	min := func(args ...int) int {
		ans := args[0]
		for i := 0; i < len(args); i++ {
			if args[i] < ans {
				ans = args[i]
			}
		}
		return ans
	}
	ans := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			h := grid[i][j]
			if h == 0 {
				continue
			}
			ans += 6*h - 2*(h-1)
			if i > 0 {
				ans -= min(h, grid[i-1][j])
			}
			if j+1 < n {
				ans -= min(h, grid[i][j+1])
			}
			if i+1 < m {
				ans -= min(h, grid[i+1][j])
			}
			if j > 0 {
				ans -= min(h, grid[i][j-1])
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
