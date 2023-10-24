package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1893_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isCovered(ranges [][]int, left int, right int) bool {
	nums := make(map[int]struct{}, 64)
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			nums[i] = struct{}{}
		}
	}
	for i := left; i <= right; i++ {
		if _, ok := nums[i]; !ok {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
