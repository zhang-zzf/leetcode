package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1184_thenSuccess(t *testing.T) {
	ans := distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2)
	assert.Equal(t, 3, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	c, sum := 0, 0
	for i := 0; i < len(distance); i++ {
		sum += distance[i]
		if i >= start && i < destination {
			c += distance[i]
		}
	}
	ans := sum - c
	if ans > c {
		ans = c
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
