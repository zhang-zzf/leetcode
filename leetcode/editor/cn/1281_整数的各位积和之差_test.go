package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1281_thenSuccess(t *testing.T) {
	assert.Equal(t, 15, subtractProductAndSum(234))
}

//leetcode submit region begin(Prohibit modification and deletion)
func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		nn := n % 10
		sum += nn
		product *= nn
		n /= 10
	}
	return product - sum
}

//leetcode submit region end(Prohibit modification and deletion)
