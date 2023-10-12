package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1684_thenSuccess(t *testing.T) {
	ans := countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"})
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countConsistentStrings(allowed string, words []string) (ans int) {
	allowedMap := make([]byte, 26)
	for _, c := range allowed {
		allowedMap[c-'a'] = 1
	}
	for _, w := range words {
		idx, n := 0, len(w)
		for ; idx < n; idx++ {
			if allowedMap[w[idx]-'a'] != 1 {
				break
			}
		}
		if idx == n {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
