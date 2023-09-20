package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1304_thenSuccess(t *testing.T) {
	ans := sumZero(6)
	sum := 0
	for _, n := range ans {
		sum += n
	}
	assert.Equal(t, 0, sum)
}

// n 个 各不相同 的整数组成的数组
func Test_givenFailedCase1_when1304_thenSuccess(t *testing.T) {
	ans := sumZero(2)
	sum := 0
	for _, n := range ans {
		sum += n
	}
	assert.Equal(t, 0, sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumZero(n int) []int {
	if n == 2 {
		return []int{-1, 1}
	}
	ans := make([]int, n)
	for i := 0; i < n-1; i++ {
		ans[i] = i
	}
	ans[n-1] = -(n - 2) * (n - 1) / 2
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
