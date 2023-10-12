package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1688_thenSuccess(t *testing.T) {
	assert.Equal(t, 6, numberOfMatches(7))
	assert.Equal(t, 13, numberOfMatches(14))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfMatches(n int) (ans int) {
	for n > 1 {
		ans += n / 2
		if n&0x01 == 1 {
			n = n/2 + 1
		} else {
			n /= 2
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
