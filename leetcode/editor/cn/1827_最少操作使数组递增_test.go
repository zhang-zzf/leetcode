package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1827_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, minOperations([]int{1, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		op := nums[i-1] + 1 - nums[i]
		if op > 0 {
			ans += op
			nums[i] = nums[i-1] + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
