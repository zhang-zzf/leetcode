package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1523_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, countOdds(3, 7))
	assert.Equal(t, 3, countOdds(3, 8))
	assert.Equal(t, 1, countOdds(8, 10))
	assert.Equal(t, 2, countOdds(8, 11))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countOdds(low int, high int) int {
	ans := (high - low + 1) / 2
	if low&0x01 == 1 && high&0x01 == 1 {
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
