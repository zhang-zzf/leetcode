package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when693_thenSuccess(t *testing.T) {
	assert.Equal(t, true, hasAlternatingBits(5))
	assert.Equal(t, false, hasAlternatingBits(7))
	assert.Equal(t, false, hasAlternatingBits(6))
}

func Test_givenEventOdd_when693_thenSuccess(t *testing.T) {
	for i := 0; i < math.MaxInt32; i++ {
		if i%2 == 0 {
			assert.Equal(t, 0, i&0x01)
		} else {
			assert.Equal(t, 1, i&0x01)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func hasAlternatingBits(n int) bool {
	a := n ^ (n >> 1)
	return a&(a+1) == 0
}

func hasAlternatingBits1(n int) bool {
	// TODO n%2 == n&0x01
	bit := n & 0x01
	for n > 0 {
		_0 := n & 0x01
		if bit != _0 {
			break
		}
		n >>= 1
		if bit == 0 {
			bit = 1
		} else {
			bit = 0
		}
	}
	return n == 0
}

//leetcode submit region end(Prohibit modification and deletion)
