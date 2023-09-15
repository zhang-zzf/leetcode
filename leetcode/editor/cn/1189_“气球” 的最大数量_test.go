package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1189_thenSuccess(t *testing.T) {
	ans := maxNumberOfBalloons("loonbalxballpoon")
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxNumberOfBalloons(text string) int {
	// 基数统计
	dp := make([]int, 26)
	for _, c := range text {
		dp[c-'a'] += 1
	}
	ans := 0
	flag := true
	for {
		for _, c := range "balloon" {
			if dp[c-'a'] -= 1; dp[c-'a'] < 0 {
				flag = false
			}
		}
		if !flag {
			break
		}
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
