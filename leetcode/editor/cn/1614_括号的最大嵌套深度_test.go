package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1614_thenSuccess(t *testing.T) {
	ans := maxDepth("(1+(2*3)+((8)/4))+1")
	assert.Equal(t, 3, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxDepth(s string) int {
	depth, maxDepth := 0, 0
	for _, c := range s {
		if c == '(' {
			depth += 1
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if c == ')' {
			depth -= 1
		}
	}
	return maxDepth
}

//leetcode submit region end(Prohibit modification and deletion)
