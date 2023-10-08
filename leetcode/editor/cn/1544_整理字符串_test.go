package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

func Test_givenNormal_when1544_thenSuccess(t *testing.T) {
	ans := makeGood("leEeetcode")
	assert.Equal(t, "leetcode", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func makeGood(s string) string {
	var ans []rune
	for _, c := range s {
		prev := c
		if unicode.IsLower(c) {
			// TODO unicode
			prev = unicode.ToUpper(c)
		} else {
			prev = unicode.ToLower(c)
		}
		n := len(ans)
		if n > 0 && ans[n-1] == prev {
			ans = ans[:n-1]
		} else {
			ans = append(ans, c)
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
