package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when405_thenSuccess(t *testing.T) {
	assert.Equal(t, "1a", toHex(26))
	assert.Equal(t, "ffffffff", toHex(-1))
}

/**
边界条件 0
*/
func Test_givenFailedCase1_when405_thenSuccess(t *testing.T) {
	assert.Equal(t, "0", toHex(0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hexTab := []string{"0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9", "a", "b", "c", "d", "e", "f"}
	mask := uint32(0x0f)
	ans := ""
	n := uint32(num)
	for n != 0 {
		lastFourBits := n & mask
		ans = hexTab[lastFourBits] + ans
		n >>= 4
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
