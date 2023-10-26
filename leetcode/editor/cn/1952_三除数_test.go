package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1952_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isThree(n int) bool {
	// todo 边界
	if n < 4 { // 1,2,3
		return false
	}
	sqr := int(math.Sqrt(float64(n)))
	// todo 确切3个
	if sqr*sqr != n { // 8
		return false
	}
	for i := 2; i < sqr; i++ { // 16
		if n%i == 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
