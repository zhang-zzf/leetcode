package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1137_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, tribonacci(4))
	assert.Equal(t, 1389537, tribonacci(25))
}

//leetcode submit region begin(Prohibit modification and deletion)
func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	n0, n1, n2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		n0, n1, n2 = n1, n2, n0+n1+n2
	}
	return n2
}

//leetcode submit region end(Prohibit modification and deletion)
