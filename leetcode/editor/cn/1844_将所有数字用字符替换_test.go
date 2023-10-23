package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1844_thenSuccess(t *testing.T) {
	assert.Equal(t, "abcdef", replaceDigits("a1c1e1"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func replaceDigits(s string) string {
	var ans []byte
	for i := 0; i < len(s); i++ {
		if i&0x01 == 0 {
			ans = append(ans, s[i])
		} else {
			ans = append(ans, s[i-1]+s[i]-'0')
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
