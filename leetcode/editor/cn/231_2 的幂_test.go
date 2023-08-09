package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when231_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isPowerOfTwo(1))
	assert.Equal(t, true, isPowerOfTwo(16))
	assert.Equal(t, false, isPowerOfTwo(15))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	pivot := 1
	for pivot < n {
		pivot <<= 1
	}
	return pivot == n
}

//leetcode submit region end(Prohibit modification and deletion)
