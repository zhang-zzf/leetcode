package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1646_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, getMaximumGenerated(7))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getMaximumGenerated(n int) (ans int) {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if i&0x01 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + dp[i/2+1]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
