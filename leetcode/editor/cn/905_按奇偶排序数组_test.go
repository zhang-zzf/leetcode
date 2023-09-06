package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when905_thenSuccess(t *testing.T) {
	assert.Equal(t, []int{4, 2, 1, 3}, sortArrayByParity([]int{3, 2, 1, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参看答案：快慢双指针， 次题和快速排序一个思路
func sortArrayByParity(nums []int) []int {
	for l, r := -1, 0; r < len(nums); r++ {
		if nums[r]%2 == 0 {
			l += 1
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return nums
}

// 左右双指针
func sortArrayByParity1(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		if nums[l]%2 == 0 {
			l += 1
			continue
		}
		if nums[r]%2 != 0 {
			r -= 1
			continue
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
