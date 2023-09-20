package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1323_thenSuccess(t *testing.T) {
	assert.Equal(t, 9969, maximum69Number(9669))
	assert.Equal(t, 99, maximum69Number(99))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximum69Number(num int) int {
	high6 := 0
	for i, n := 1, num; n > 0; i++ {
		mod := n % 10
		if mod == 6 {
			high6 = i
		}
		n /= 10
	}
	if high6 > 0 {
		num += int(3 * math.Pow10(high6-1))
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
