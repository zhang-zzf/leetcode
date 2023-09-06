package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when917_thenSuccess(t *testing.T) {
	ans := reverseOnlyLetters("Test1ng-Leet=code-Q!")
	assert.Equal(t, "Qedo1ct-eeLg=ntse-T!", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(s string) string {
	isChar := func(c byte) bool {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			return true
		}
		return false
	}
	bytes := []byte(s)
	for l, r := 0, len(bytes)-1; l < r; {
		if !isChar(bytes[l]) {
			l += 1
			continue
		}
		if !isChar(bytes[r]) {
			r -= 1
			continue
		}
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l, r = l+1, r-1
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
