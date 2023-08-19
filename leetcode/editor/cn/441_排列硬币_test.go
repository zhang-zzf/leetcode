package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when441_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, arrangeCoins(8))
	assert.Equal(t, 2, arrangeCoins(5))
}

/**
3
*/
func Test_givenFailedCase1_when441_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, arrangeCoins(3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	// 值域2分
	ans := 1
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		sum := mid * (mid + 1) / 2
		if sum <= n {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
