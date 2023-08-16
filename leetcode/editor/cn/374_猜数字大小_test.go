package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when374_thenSuccess(t *testing.T) {
	assert.Equal(t, 6, guessNumber(10))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right, ans := 1, n, 1
	for left <= right {
		mid := left + (right-left)>>1
		pivot := guess(mid)
		if pivot == 0 {
			ans = mid
			break
		} else if pivot < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
func guess(num int) int {
	if num == 6 {
		return 0
	} else if num > 6 {
		return -1
	} else {
		return 1
	}
}
