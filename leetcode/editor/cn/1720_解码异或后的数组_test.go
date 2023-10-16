package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1720_thenSuccess(t *testing.T) {
	assert.Equal(t, []int{1, 0, 2, 1}, decode([]int{1, 2, 3}, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func decode(encoded []int, first int) []int {
	n := len(encoded)
	ans := make([]int, n+1)
	ans[0] = first
	for i := 0; i < n; i++ {
		ans[i+1] = encoded[i] ^ ans[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
