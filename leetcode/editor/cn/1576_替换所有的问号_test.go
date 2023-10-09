package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1576_thenSuccess(t *testing.T) {
	assert.Equal(t, "ubvabaw", modifyString("ubv???w"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func modifyString(s string) string {
	bytes := make([]byte, len(s))
	for i, c := range []byte(s) {
		bytes[i] = c
		if c != '?' {
			continue
		}
		predecessor, successor := byte(0), byte(0)
		if i > 0 {
			predecessor = bytes[i-1]
		}
		if i+1 < len(s) {
			successor = s[i+1]
		}
		for candidate := byte('a'); candidate < 'z'; candidate++ {
			if candidate != predecessor && candidate != successor {
				bytes[i] = candidate
				break
			}
		}
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
