package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when724_thenSuccess(t *testing.T) {
	assert.Equal(t, 0, pivotIndex([]int{2, 1, -1}))
	assert.Equal(t, 3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(t, -1, pivotIndex([]int{1, 2, 3}))
}

func Test_givenFailedCase1_when724_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, pivotIndex([]int{-1, -1, 0, 0, -1, -1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func pivotIndex(nums []int) int {
	ans := -1
	// prefix sum
	n := len(nums)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	for i := 0; i < n; i++ {
		if sum[i]*2 == sum[n]-nums[i] {
			ans = i
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
