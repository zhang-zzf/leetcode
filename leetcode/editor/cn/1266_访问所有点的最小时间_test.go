package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1266_thenSuccess(t *testing.T) {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	assert.Equal(t, 7, minTimeToVisitAllPoints(points))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := 1; i < len(points); i++ {
		ans += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}
	return ans
}

func minTimeToVisitAllPoints1(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		x, y := points[i][0]-points[i-1][0], points[i][1]-points[i-1][1]
		// todo 绝对值
		x, y = int(math.Abs(float64(x))), int(math.Abs(float64(y)))
		xy := x
		if xy < y {
			xy = y
		}
		ans += xy
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
