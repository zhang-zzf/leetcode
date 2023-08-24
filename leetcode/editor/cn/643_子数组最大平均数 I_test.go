package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when643_thenSuccess(t *testing.T) {
	ans := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	assert.Equal(t, 12.75, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	// 滑动窗口
	l, r, sum := 0, 0, 0
	for ; r < k; r++ {
		sum += nums[r]
	}
	maxSum := sum
	for r < len(nums) {
		sum += nums[r]
		sum -= nums[l]
		l, r = l+1, r+1
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}

func findMaxAverage1(nums []int, k int) float64 {
	// 前缀和
	n := len(nums)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	ans := math.MinInt
	for i := k - 1; i < n; i++ {
		sum := preSum[i+1] - preSum[i+1-k]
		if sum > ans {
			ans = sum
		}
	}
	return float64(ans) / float64(k)
}

//leetcode submit region end(Prohibit modification and deletion)
