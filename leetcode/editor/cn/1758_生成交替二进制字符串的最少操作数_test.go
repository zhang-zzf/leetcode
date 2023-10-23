package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1758_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, minOperations1758("0100"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations1758(s string) int {
	_0, _1 := 0, 0
	for _, c := range s {
		if c == '0' {
			_0, _1 = _1, _0+1
		} else {
			_0, _1 = _1+1, _0
		}
	}
	if _0 < _1 {
		return _0
	}
	return _1
}

//leetcode submit region end(Prohibit modification and deletion)
