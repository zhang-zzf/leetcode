package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1446_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, maxPower("abbcccddddeeeeedcba"))
	assert.Equal(t, 2, maxPower("leetcode"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxPower(s string) int {
	cnt, max := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt += 1
		} else {
			cnt = 1
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
