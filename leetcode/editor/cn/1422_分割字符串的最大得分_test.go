package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1422_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, maxScore("1111"))
	assert.Equal(t, 0, maxScore("10"))
	assert.Equal(t, 2, maxScore("01"))
	assert.Equal(t, 5, maxScore("00111"))
	assert.Equal(t, 5, maxScore("011101"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxScore(s string) int {
	l, r, ans, n := 0, 0, 0, len(s)
	if s[0] == '0' {
		l += 1
	}
	for _, c := range s[1:] {
		if c == '1' {
			r += 1
		}
	}
	ans = l + r
	for _, c := range s[1 : n-1] {
		if c == '0' {
			l += 1
		} else if c == '1' {
			r -= 1
		}
		if l+r > ans {
			ans = l + r
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
