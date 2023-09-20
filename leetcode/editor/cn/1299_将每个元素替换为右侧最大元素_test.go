package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1299_thenSuccess(t *testing.T) {
	ans := replaceElements([]int{17, 18, 5, 4, 6, 1})
	assert.Equal(t, []int{18, 6, 6, 6, 1, -1}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func replaceElements(arr []int) []int {
	rightMax := -1
	n := len(arr)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[i] = rightMax
		if arr[i] > rightMax {
			rightMax = arr[i]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
