package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when290_thenSuccess(t *testing.T) {
	assert.Equal(t, true, wordPattern("abba", "d a a d"))
	assert.Equal(t, false, wordPattern("abba", "d d d d"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	ans := true
	p2w, w2p := map[uint8]string{}, map[string]uint8{}
	for i := 0; i < len(pattern); i++ {
		p, w := pattern[i], words[i]
		// 判断 map 中是否存在 key
		m1, e1 := p2w[p]
		m2, e2 := w2p[w]
		if !e1 && !e2 {
			p2w[p] = w
			w2p[w] = p
		} else if m1 != w || m2 != p {
			ans = false
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
