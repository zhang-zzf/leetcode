package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when804_thenSuccess(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	assert.Equal(t, 2, uniqueMorseRepresentations(words))
	assert.Equal(t, 1, uniqueMorseRepresentations([]string{"a"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morse := map[string]struct{}{}
	for _, w := range words {
		m := strings.Builder{}
		for _, c := range []byte(w) {
			m.WriteString(codes[c-'a'])
		}
		morse[m.String()] = struct{}{}
	}
	return len(morse)
}

//leetcode submit region end(Prohibit modification and deletion)
