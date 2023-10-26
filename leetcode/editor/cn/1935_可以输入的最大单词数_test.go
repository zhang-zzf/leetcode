package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1935_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, canBeTypedWords("hello world", "ad"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	ans := len(words)
	for _, w := range words {
		if strings.ContainsAny(w, brokenLetters) {
			ans -= 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
