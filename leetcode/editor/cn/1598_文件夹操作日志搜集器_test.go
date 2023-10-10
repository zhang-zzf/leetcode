package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1598_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(logs []string) int {
	depth := 0
	for _, log := range logs {
		switch log {
		case "../":
			depth -= 1
			if depth < 0 {
				depth = 0
			}
		case "./":
		default:
			depth += 1
		}
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)
