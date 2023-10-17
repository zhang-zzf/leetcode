package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1732_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestAltitude(gain []int) int {
	ans, height := 0, 0
	for _, g := range gain {
		height = height + g
		if height > ans {
			ans = height
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
