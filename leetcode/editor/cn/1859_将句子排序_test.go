package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1859_thenSuccess(t *testing.T) {
	ans := sortSentence("is2 sentence4 This1 a3")
	assert.Equal(t, "This is a sentence", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortSentence(s string) string {
	words := strings.Split(s, " ")
	mapping := make([]string, len(words))
	for _, w := range words {
		last := len(w) - 1
		mapping[w[last]-'1'] = w[:last]
	}
	return strings.Join(mapping, " ")
}

//leetcode submit region end(Prohibit modification and deletion)
