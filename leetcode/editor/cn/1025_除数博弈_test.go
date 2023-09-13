package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1025_thenSuccess(t *testing.T) {
	assert.Equal(t, false, divisorGame(1))
	assert.Equal(t, true, divisorGame(2))
	assert.Equal(t, false, divisorGame(3))
	assert.Equal(t, true, divisorGame(4))
	assert.Equal(t, false, divisorGame(5))
	assert.Equal(t, true, divisorGame(6))
}

//leetcode submit region begin(Prohibit modification and deletion)
func divisorGame(n int) bool {
	// 动态规划
	dp := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if dp[i-1] == false {
			dp[i] = true
		} else {
			for j := 2; j*j <= i; j++ {
				if i%j == 0 && (dp[i-j] == false || dp[i-i/j] == false) {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
