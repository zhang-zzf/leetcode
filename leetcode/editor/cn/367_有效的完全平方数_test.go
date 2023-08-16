package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when367_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isPerfectSquare(16))
	assert.Equal(t, false, isPerfectSquare(12))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	// TODO Divide and Conquer
	left, right := 1, num
	ans := 1
	for left <= right {
		mid := left + (right-left)>>1
		square := mid * mid
		if square == num {
			ans = mid
			break
		} else if square > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans*ans == num
}

//leetcode submit region end(Prohibit modification and deletion)
