package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1678_thenSuccess(t *testing.T) {
	assert.Equal(t, "alGalooG", interpret("(al)G(al)()()G"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func interpret(command string) string {
	var ans []byte
	var stack []byte
	for _, c := range []byte(command) {
		switch c {
		case 'G':
			ans = append(ans, c)
		case ')':
			if len(stack) == 1 {
				ans = append(ans, 'o')
			} else {
				ans = append(ans, stack[1:]...)
			}
			stack = []byte{}
		default:
			stack = append(stack, c)
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
