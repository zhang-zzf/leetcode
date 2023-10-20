package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1800_thenSuccess(t *testing.T) {
	assert.Equal(t, 33, maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxAscendingSum(nums []int) int {
	sum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

//leetcode submit region end(Prohibit modification and deletion)
