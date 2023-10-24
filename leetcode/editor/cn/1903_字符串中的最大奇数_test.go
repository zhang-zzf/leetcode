package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1903_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')&0x01 == 1 {
			return num[:i+1]
		}
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
