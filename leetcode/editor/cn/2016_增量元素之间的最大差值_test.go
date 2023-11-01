package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2016_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, maximumDifference([]int{7, 1, 5, 4}))
	assert.Equal(t, -1, maximumDifference([]int{1, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumDifference(nums []int) int {
	ans := 0
	n := len(nums)
	for r := 1; r < n; r++ {
		for l := 0; l < r; l++ {
			if diff := nums[r] - nums[l]; diff > ans {
				ans = diff
			}
		}
	}
	if ans == 0 {
		return -1
	}
	return ans
}

// todo 题意理解错误，理解成了上升趋势下的最大值
func maximumDifference2(nums []int) int {
	ans := -1
	low, high := 0, 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff <= 0 {
			low = i
		} else if diff > 0 {
			high = i
		}
		if n := nums[high] - nums[low]; low < high && n > ans {
			ans = n
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
