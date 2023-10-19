package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

func Test_givenNormal_when1763_thenSuccess(t *testing.T) {
	assert.Equal(t, "aAa", longestNiceSubstring("YazaAay"))
	assert.Equal(t, "Bb", longestNiceSubstring("Bb"))
	assert.Equal(t, "", longestNiceSubstring("Aba"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案，暴力破解
func longestNiceSubstring(s string) string {
	n := len(s)
	maxLng, start := 0, 0
	for i := 0; i < n; i++ {
		lower, upper := 0, 0
		for j := i; j < n; j++ {
			c := rune(s[j])
			if unicode.IsLower(c) {
				// TODO 相应位置置1
				lower |= 1 << (c - 'a')
			} else {
				upper |= 1 << (c - 'A')
			}
			if lower == upper && j-i+1 > maxLng {
				maxLng = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLng]
}

//leetcode submit region end(Prohibit modification and deletion)
