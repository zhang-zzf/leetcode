package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when190_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseBits(num uint32) uint32 {
	const (
		M8 = 0x00ff00ff
		M4 = 0x0f0f0f0f
		M2 = 0x33333333
		M1 = 0x55555555
	)
	// 使用括号后，不用考虑符号优先级
	num = ((num >> 1) & M1) | ((num & M1) << 1)
	num = ((num >> 2) & M2) | ((num & M2) << 2)
	num = ((num >> 4) & M4) | ((num & M4) << 4)
	num = ((num >> 8) & M8) | ((num & M8) << 8)
	num = num>>16 | num<<16
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
