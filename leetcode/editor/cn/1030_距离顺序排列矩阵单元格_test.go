package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1030_thenSuccess(t *testing.T) {
	ans := allCellsDistOrder(2, 2, 0, 1)
	assert.Equal(t, [][]int{{0, 1}, {1, 1}, {0, 0}, {1, 0}}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	var ans [][]int
	// bfs
	// TODO
	dp := make([][]int, 1, rows*cols)
	dp[rCenter][cCenter] = math.MaxInt32
	queue := [][]int{{rCenter, cCenter}}
	directs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		ans = append(ans, cell)
		for _, d := range directs {
			r, c := cell[0]+d[0], cell[1]+d[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && dp[r][c] == 0 {
				dp[r][c] = 1
				queue = append(queue, []int{r, c})
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
