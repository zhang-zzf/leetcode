package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1437_thenSuccess(t *testing.T) {
	assert.Equal(t, true, kLengthApart([]int{1, 1, 1, 1, 1}, 0))
	assert.Equal(t, false, kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

func Test_givenFailedCase1_when1437_thenSuccess(t *testing.T) {
	assert.Equal(t, true, kLengthApart([]int{0, 1, 0, 0, 1, 0, 0, 1}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func kLengthApart(nums []int, k int) bool {
	zeroCnt, oneCnt := 0, 0
	for _, n := range nums {
		if n == 0 {
			zeroCnt += 1
		} else {
			if zeroCnt < k && oneCnt > 0 {
				return false
			}
			zeroCnt, oneCnt = 0, oneCnt+1
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
