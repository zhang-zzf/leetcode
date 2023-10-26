package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1957_thenSuccess(t *testing.T) {
	assert.Equal(t, "aabaa", makeFancyString("aaabaaaa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func makeFancyString(s string) string {
	ans := []byte{s[0]}
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt += 1
			if cnt >= 3 {
				continue
			}
		} else {
			cnt = 1
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
