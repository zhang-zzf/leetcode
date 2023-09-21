package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1342_thenSuccess(t *testing.T) {
	assert.Equal(t, 6, numberOfSteps(14))
	assert.Equal(t, 4, numberOfSteps(8))
	assert.Equal(t, 12, numberOfSteps(123))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfSteps(num int) int {
	ans := 0
	for num > 0 {
		if num&0x01 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
