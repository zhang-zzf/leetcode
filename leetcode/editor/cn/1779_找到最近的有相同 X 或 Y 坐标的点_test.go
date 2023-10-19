package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1779_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func nearestValidPoint(x int, y int, points [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	ans, distance := -1, math.MaxInt
	for idx, ptr := range points {
		d1, d2 := abs(x-ptr[0]), abs(y-ptr[1])
		if (d1 == 0 || d2 == 0) && (d1+d2) < distance {
			distance = d1 + d2
			ans = idx
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
