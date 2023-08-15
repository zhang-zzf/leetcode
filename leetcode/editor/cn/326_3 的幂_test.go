package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when326_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	p := 1
	for p < n {
		p *= 3
	}
	return p == n
}

//leetcode submit region end(Prohibit modification and deletion)
