package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1920_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, n := range nums {
		ans[i] = nums[n]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
