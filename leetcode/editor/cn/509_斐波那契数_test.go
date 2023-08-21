package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when509_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, fib(1))
	assert.Equal(t, 1, fib(2))
	assert.Equal(t, 2, fib(3))
	assert.Equal(t, 3, fib(4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	// dp
	if n < 2 {
		return n
	}
	n2, n1 := 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n1
}

//leetcode submit region end(Prohibit modification and deletion)
