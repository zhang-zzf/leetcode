package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1491_thenSuccess(t *testing.T) {
	ans := average([]int{6000, 5000, 4000, 3000, 2000, 1000})
	assert.Equal(t, float64(3500), ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func average(salary []int) float64 {
	min, max, sum := math.MaxInt, math.MinInt, 0
	for _, n := range salary {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}

//leetcode submit region end(Prohibit modification and deletion)
