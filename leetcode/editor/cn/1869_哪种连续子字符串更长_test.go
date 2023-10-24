package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1869_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkZeroOnes(s string) bool {
	_0, _1, _0m, _1m := 0, 0, 0, 0
	for _, c := range s {
		if c == '1' {
			_0 = 0
			_1 += 1
			if _1 > _1m {
				_1m = _1
			}
		} else {
			_1 = 0
			_0 += 1
			if _0 > _0m {
				_0m = _0
			}
		}
	}
	return _1m > _0m
}

//leetcode submit region end(Prohibit modification and deletion)
