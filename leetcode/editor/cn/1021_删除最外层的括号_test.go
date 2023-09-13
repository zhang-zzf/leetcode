package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1021_thenSuccess(t *testing.T) {
	ans := removeOuterParentheses("(()())(())")
	assert.Equal(t, "()()()", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeOuterParentheses(s string) string {
	ans := strings.Builder{}
	var bytes []byte
	leftCnt := 0
	for _, c := range []byte(s) {
		bytes = append(bytes, c)
		if c == '(' {
			leftCnt += 1
		} else {
			leftCnt -= 1
		}
		if leftCnt == 0 {
			bytes = bytes[1 : len(bytes)-1]
			ans.Write(bytes)
			bytes = []byte{}
		}
	}
	return ans.String()
}

//leetcode submit region end(Prohibit modification and deletion)
