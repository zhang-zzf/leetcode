package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when495_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, findPoisonedDuration([]int{1, 2}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		period := timeSeries[i+1] - timeSeries[i]
		if period > duration {
			period = duration
		}
		ans += period
	}
	// 最后一次 duration
	ans += duration
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
