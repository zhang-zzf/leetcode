package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when824_thenSuccess(t *testing.T) {
	a := toGoatLatin("I speak Goat Latin")
	assert.Equal(t, "Imaa peaksmaaa oatGmaaaa atinLmaaaaa", a)
	b := toGoatLatin("The quick brown fox jumped over the lazy dog")
	assert.Equal(t, "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa", b)
}

//leetcode submit region begin(Prohibit modification and deletion)
func toGoatLatin(sentence string) string {
	aeiou := map[byte]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	ans := strings.Builder{}
	sentence = sentence + " "
	n, wi := len(sentence), 1
	for r, l := 0, 0; r < n; r++ {
		if sentence[r] != ' ' {
			continue
		}
		// sentence[r] == ' '
		if sentence[l] != ' ' {
			// [l,r) is a word
			if _, ok := aeiou[sentence[l]]; ok {
				for i := l; i < r; i++ {
					ans.WriteByte(sentence[i])
				}
			} else {
				for i := l + 1; i < r; i++ {
					ans.WriteByte(sentence[i])
				}
				ans.WriteByte(sentence[l])
			}
			ans.WriteString("ma")
			for i := 0; i < wi; i++ {
				ans.WriteByte('a')
			}
			wi += 1
		}
		if r != n-1 {
			ans.WriteByte(' ')
		}
		l = r + 1
	}
	return ans.String()
}

//leetcode submit region end(Prohibit modification and deletion)
