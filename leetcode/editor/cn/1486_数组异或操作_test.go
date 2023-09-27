package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1486_thenSuccess(t *testing.T) {
	ans := xorOperation(10, 5)
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func xorOperation(n int, start int) int {
	ans := start
	for i := 1; i < n; i++ {
		ans ^= start + 2*i
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
