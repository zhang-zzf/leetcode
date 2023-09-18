package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1221_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, balancedStringSplit("RLRRRLLRLL"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func balancedStringSplit(s string) (ans int) {
	rc, lc := 0, 0
	for _, c := range s {
		if c == 'R' {
			rc += 1
		} else {
			lc += 1
		}
		if lc == rc {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
