package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when771_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, numJewelsInStones("aA", "aAAbbbb"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numJewelsInStones(jewels string, stones string) int {
	dp := make([]int, 52)
	for _, c := range []byte(stones) {
		if c < 'a' {
			dp[c-'A'] += 1
		} else {
			dp[c-'a'+26] += 1
		}
	}
	ans := 0
	for _, c := range []byte(jewels) {
		if c < 'a' {
			ans += dp[c-'A']
		} else {
			ans += dp[c-'a'+26]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
