package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1413_thenSuccess(t *testing.T) {
	ans := minStartValue([]int{-3, 2, -3, 4, 2})
	assert.Equal(t, 5, ans)
}

func Test_givenFailedCase1_when1413_thenSuccess(t *testing.T) {
	ans := minStartValue([]int{1, 2})
	assert.Equal(t, 1, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minStartValue(nums []int) int {
	preSum, minSum := 0, math.MaxInt
	for _, n := range nums {
		preSum += n
		if preSum < minSum {
			minSum = preSum
		}
	}
	if minSum > 0 {
		return 1
	} else {
		return -minSum + 1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
