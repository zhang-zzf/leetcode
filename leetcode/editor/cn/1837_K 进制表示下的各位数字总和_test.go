package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1837_thenSuccess(t *testing.T) {
	assert.Equal(t, 9, sumBase(34, 6))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumBase(n int, k int) int {
	ans := 0
	for n > 0 {
		ans += n % k
		n /= k
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
