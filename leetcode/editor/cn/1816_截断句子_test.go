package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1816_thenSuccess(t *testing.T) {
	ans := truncateSentence("Hello how are you Contestant", 4)
	assert.Equal(t, "Hello how are you", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func truncateSentence(s string, k int) string {
	var ans []byte
	for _, c := range []byte(s) {
		if c == ' ' {
			k -= 1
		}
		if k <= 0 {
			break
		}
		ans = append(ans, c)
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
