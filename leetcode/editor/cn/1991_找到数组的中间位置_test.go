package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1991_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, findMiddleIndex([]int{2, 3, -1, 8, 4}))
	assert.Equal(t, -1, findMiddleIndex([]int{2, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMiddleIndex(nums []int) int {
	n := len(nums)
	leftSum, sum := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	for i := 0; i < n; i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
