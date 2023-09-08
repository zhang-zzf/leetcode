package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when976_thenSuccess(t *testing.T) {
	assert.Equal(t, 7, largestPerimeter([]int{1, 2, 1, 2, 3, 10}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			ans = nums[i-1] + nums[i-2] + nums[i]
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
