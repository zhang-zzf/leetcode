package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1704_thenSuccess(t *testing.T) {
	assert.Equal(t, true, halvesAreAlike("book"))
	assert.Equal(t, false, halvesAreAlike("textbook"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func halvesAreAlike(s string) bool {
	vowels := make([]byte, 128)
	for _, v := range "aeiouAEIOU" {
		vowels[v] = 1
	}
	lc, rc := 0, 0
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if vowels[s[l]] == 1 {
			lc += 1
		}
		if vowels[s[r]] == 1 {
			rc += 1
		}
	}
	return lc == rc
}

//leetcode submit region end(Prohibit modification and deletion)
