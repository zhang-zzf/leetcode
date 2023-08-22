package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when598_thenSuccess(t *testing.T) {
	ans := maxCount(3, 3, [][]int{{2, 2}, {3, 3}})
	assert.Equal(t, 4, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		if op[0] < m {
			m = op[0]
		}
		if op[1] < n {
			n = op[1]
		}
	}
	return m * n
}

//leetcode submit region end(Prohibit modification and deletion)
