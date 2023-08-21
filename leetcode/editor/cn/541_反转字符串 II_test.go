package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when541_thenSuccess(t *testing.T) {
	assert.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	assert.Equal(t, "bacd", reverseStr("abcd", 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	bytes := []byte(s)
	last := len(bytes) - 1
	for i := 0; i < len(bytes); i += 2 * k {
		l, r := i, i+k-1
		if r > last {
			r = last
		}
		for ; l < r; l, r = l+1, r-1 {
			bytes[l], bytes[r] = bytes[r], bytes[l]
		}
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
