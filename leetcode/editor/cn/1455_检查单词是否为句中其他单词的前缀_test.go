package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1455_thenSuccess(t *testing.T) {
	// when
	prefixOfWord := isPrefixOfWord("i love eating burger", "burg")
	// then
	assert.Equal(t, 4, prefixOfWord, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPrefixOfWord(sentence string, searchWord string) int {
	ans := -1
	words := strings.Split(sentence, " ")
	for idx, word := range words {
		if strings.HasPrefix(word, searchWord) {
			ans = idx + 1
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
