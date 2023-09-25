package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1399_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, countLargestGroup(24))
	assert.Equal(t, 4, countLargestGroup(13))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countLargestGroup(n int) int {
	// 数位和最大值 9999 = 36
	numBitSum := make([]int, 37)
	bitSum := func(a int) int {
		ans := 0
		for a > 0 {
			ans += a % 10
			a /= 10
		}
		return ans
	}
	for i := 1; i <= n; i++ {
		numBitSum[bitSum(i)] += 1
	}
	ans, max := 0, 0
	for _, n := range numBitSum {
		if n > max {
			max = n
			ans = 1
		} else if n == max {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
