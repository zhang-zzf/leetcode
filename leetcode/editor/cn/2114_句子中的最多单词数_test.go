package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when2114_thenSuccess(t *testing.T) {
	ans := mostWordsFound([]string{"alice and bob love leetcode",
		"i think so too",
		"this is great thanks very much"})
	assert.Equal(t, 6, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func mostWordsFound(sentences []string) int {
	ans := 0
	for _, s := range sentences {
		//wc := 1
		//for _, c := range s {
		//	if c == ' ' {
		//		wc += 1
		//	}
		//}
		wc := strings.Count(s, " ") + 1
		if wc > ans {
			ans = wc
		}
	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
