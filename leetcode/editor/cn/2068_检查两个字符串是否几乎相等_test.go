package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2068_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkAlmostEquivalent(word1 string, word2 string) bool {
	cnt := make([]int, 26)
	for _, c := range word1 {
		cnt[c-'a'] += 1
	}
	for _, c := range word2 {
		cnt[c-'a'] -= 1
	}
	for _, c := range cnt {
		if c > 3 || c < -3 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
