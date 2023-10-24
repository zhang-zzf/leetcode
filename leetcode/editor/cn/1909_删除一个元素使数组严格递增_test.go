package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1909_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canBeIncreasing([]int{1, 2, 10, 5, 7}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canBeIncreasing(nums []int) bool {
	cnt := 0
	stack := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		ls := len(stack)
		if nums[i] > stack[ls-1] {
			stack = append(stack, nums[i])
			continue
		}
		// nums[i] <= stack[len(stack)-1]
		// 方式1 删除 stack[len(stack)-1]
		if (ls > 1 && stack[ls-2] < nums[i]) || ls == 1 {
			stack = stack[:ls-1]
			stack = append(stack, nums[i])
		} else {
			// 方式2 删除 nums[i]
		}
		if cnt > 0 {
			return false
		}
		cnt += 1
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
