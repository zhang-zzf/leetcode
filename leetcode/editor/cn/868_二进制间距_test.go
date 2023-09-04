package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when868_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, binaryGap(22))
}

//leetcode submit region begin(Prohibit modification and deletion)
func binaryGap(n int) int {
	i, ans, left := 0, 0, 33
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			if i-left > ans {
				ans = i - left
			}
			left = i
		}
		i += 1
	}
	return ans
}
func binaryGap1(n int) int {
	var ans []byte
	for n > 0 {
		mod := n % 2
		ans = append(ans, byte(mod))
		n /= 2
	}
	left, max := 33, 0
	for idx, b := range ans {
		if b == 1 {
			if idx-left > max {
				max = idx - left
			}
			left = idx
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
