package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1592_thenSuccess(t *testing.T) {
	assert.Equal(t, "this   is   a   sentence", reorderSpaces("  this   is  a sentence "))
	assert.Equal(t, "practice   makes   perfect ", reorderSpaces(" practice   makes   perfect"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 字符串操作
func reorderSpaces(text string) string {
	words := strings.Fields(text)
	spaceCnt := strings.Count(text, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", spaceCnt)
	} else {
		return strings.Join(words, strings.Repeat(" ", spaceCnt/lw)) + strings.Repeat(" ", spaceCnt%lw)
	}
}

func reorderSpaces1(text string) string {
	word, wc, space := 0, 0, 0
	for _, c := range text {
		if c == ' ' {
			space += 1
			wc = 0
		} else {
			if wc == 0 {
				word += 1
			}
			wc = 1
		}
	}
	tail := space
	if word > 1 {
		tail, space = space%(word-1), space/(word-1)
	}
	wc = 0
	var ans []byte
	for _, c := range []byte(text) {
		if c != ' ' {
			ans = append(ans, c)
			wc = 1
		} else {
			if wc == 1 && word > 1 {
				for i := 0; i < space; i++ {
					ans = append(ans, ' ')
				}
				word -= 1
			}
			wc = 0
		}
	}
	for ; tail > 0; tail-- {
		ans = append(ans, ' ')
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
