package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when844_thenSuccess(t *testing.T) {
	assert.Equal(t, true, backspaceCompare("ab#c", "ad#c"))
}

func Test_givenFailedCase1_when844_thenSuccess(t *testing.T) {
	assert.Equal(t, false, backspaceCompare("a#c", "b"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare(s string, t string) bool {
	backspace := func(str string) string {
		var ans []rune
		for _, c := range str {
			if c != '#' {
				ans = append(ans, c)
				continue
			}
			if len(ans) != 0 {
				ans = ans[:len(ans)-1]
			}
		}
		return string(ans)
	}
	return backspace(s) == backspace(t)
}

//leetcode submit region end(Prohibit modification and deletion)
