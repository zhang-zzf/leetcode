package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1941_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func areOccurrencesEqual(s string) bool {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a'] += 1
	}
	pivot := 0
	for _, c := range cnt {
		if c == 0 {
			continue
		}
		if pivot == 0 {
			pivot = c
			continue
		}
		if c != pivot {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
