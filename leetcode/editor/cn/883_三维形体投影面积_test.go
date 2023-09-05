package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when883_thenSuccess(t *testing.T) {
	assert.Equal(t, 17, projectionArea([][]int{{1, 2}, {3, 4}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func projectionArea(grid [][]int) int {
	// x, y
	area := 0
	for _, row := range grid {
		rowMax := 0
		for _, n := range row {
			if n != 0 {
				area += 1 // x
			}
			if n > rowMax {
				rowMax = n
			}
		}
		area += rowMax
	}
	// z
	m, n := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		columnMax := 0
		for j := 0; j < m; j++ {
			if grid[j][i] > columnMax {
				columnMax = grid[j][i]
			}
		}
		area += columnMax
	}
	return area
}

//leetcode submit region end(Prohibit modification and deletion)
