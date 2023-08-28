package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when704_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			ans = mid
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
