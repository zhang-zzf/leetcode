package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when461_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, hammingDistance(1, 4))
	assert.Equal(t, 1, hammingDistance(1, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func hammingDistance(x int, y int) int {
	xor := x ^ y
	ans := 0
	for xor != 0 {
		xor &= xor - 1
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
