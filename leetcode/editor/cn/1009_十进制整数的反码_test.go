package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1009_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, bitwiseComplement(5))
}

func Test_givenFailedCase1_when1009_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, bitwiseComplement(0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	ans := 0
	for k := 0; n > 0; k++ {
		if n&0x01 == 0 {
			ans += 1 << k
		}
		n >>= 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
