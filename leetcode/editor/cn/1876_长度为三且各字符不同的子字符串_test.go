package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1876_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, countGoodSubstrings("aababcabc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countGoodSubstrings(s string) int {
	ans := 0
	n := len(s)
	for i := 0; i <= n-3; i++ {
		cnt := map[byte]struct{}{}
		for j := i; j < i+3; j++ {
			cnt[s[j]] = struct{}{}
		}
		if len(cnt) == 3 {
			ans += 1
		}
	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
