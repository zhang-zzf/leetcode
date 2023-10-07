package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1518_thenSuccess(t *testing.T) {
	ans := numWaterBottles(15, 4)
	assert.Equal(t, 19, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange {
		numBottles -= numExchange
		numBottles += 1
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
