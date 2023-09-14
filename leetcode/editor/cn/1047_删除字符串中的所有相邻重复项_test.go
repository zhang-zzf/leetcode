package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1047_thenSuccess(t *testing.T) {
	assert.Equal(t, "ca", removeDuplicates("abbaca"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(s string) string {
	// Stack
	var stack []byte
	for _, c := range []byte(s) {
		n := len(stack)
		if n > 0 && stack[n-1] == c {
			stack = stack[0 : n-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

//leetcode submit region end(Prohibit modification and deletion)
