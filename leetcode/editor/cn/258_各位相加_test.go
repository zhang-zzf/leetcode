package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when258_thenSuccess(t *testing.T) {
	assert.Equal(t, 0, addDigits(0))
	assert.Equal(t, 1, addDigits(1))
	assert.Equal(t, 2, addDigits(38))
}

//leetcode submit region begin(Prohibit modification and deletion)
func addDigits(num int) int {
	return (num-1)%9 + 1
}

//leetcode submit region end(Prohibit modification and deletion)
