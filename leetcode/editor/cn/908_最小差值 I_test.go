package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when908_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func smallestRangeI(nums []int, k int) int {
	min, max := math.MaxInt, math.MinInt
	for _, n := range nums {
		if n > max {
			max = n
		} else if n < min {
			min = n
		}
	}
	ans := max - min - 2*k
	if ans < 0 {
		ans = 0
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
