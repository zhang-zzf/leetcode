package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when2099_thenSuccess(t *testing.T) {
	ans := maxSubsequence([]int{-1, -2, 3, 4}, 3)
	assert.Equal(t, []int{-1, 3, 4}, ans)
}

// 不能忽略数组的有序型
func Test_givenFailedCase1_when2099_thenSuccess(t *testing.T) {
	ans := maxSubsequence([]int{50, -75}, 2)
	assert.Equal(t, []int{50, -75}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubsequence(nums []int, k int) []int {
	numAndIdx := make([][2]int, len(nums))
	for i, n := range nums {
		numAndIdx[i] = [2]int{i, n}
	}
	sort.Slice(numAndIdx, func(i, j int) bool { return numAndIdx[i][1] < numAndIdx[j][1] })
	max := numAndIdx[len(nums)-k:]
	sort.Slice(max, func(i, j int) bool { return max[i][0] < max[j][0] })
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = max[i][1]
	}
	return ans
}

func maxSubsequence1(nums []int, k int) []int {
	// sort
	sort.Ints(nums)
	return nums[len(nums)-k:]
}

//leetcode submit region end(Prohibit modification and deletion)
