package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when342_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfFour(n int) bool {
	power := 1
	for power < n {
		power *= 4
	}
	return power == n
}

//leetcode submit region end(Prohibit modification and deletion)
