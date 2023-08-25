package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when680_thenSuccess(t *testing.T) {
	assert.Equal(t, true, validPalindrome("aba"))
	assert.Equal(t, false, validPalindrome("abc"))
	assert.Equal(t, true, validPalindrome("abca"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	ans := true
	for l < r {
		if s[l] == s[r] {
			l, r = l+1, r-1
			continue
		}
		left, right := true, true
		for i, j := l+1, r; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				right = false
				break
			}
		}
		for i, j := l, r-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				left = false
				break
			}
		}
		ans = left || right
		break
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
