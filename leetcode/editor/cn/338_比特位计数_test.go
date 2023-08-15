package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when338_thenSuccess(t *testing.T) {
	bits := countBits(5)
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2}, bits)
}

/**
边界条件
n ==0 时 ans[1] index out of range
*/
func Test_givenFailedCase1_when338_thenSuccess(t *testing.T) {
	bits := countBits(0)
	assert.Equal(t, []int{0}, bits)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countBits(n int) []int {
	// 动态规划
	// 7 = 4 + 3
	// 9 = 8 + 1
	// 10 = 8 + 2
	ans := make([]int, n+1)
	power := 1
	for i := 1; i <= n; i++ {
		if i == power {
			power *= 2
			ans[i] = 1
		} else {
			lastPower := power / 2
			ans[i] = ans[lastPower] + ans[i-lastPower]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
