package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1784_thenSuccess(t *testing.T) {
	assert.Equal(t, true, checkOnesSegment("110"))
	assert.Equal(t, false, checkOnesSegment("1101"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// todo 参考答案
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func checkOnesSegment1(s string) bool {
	_0 := false
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			_0 = true
		} else if _0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
