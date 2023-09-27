package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1480_thenSuccess(t *testing.T) {
	ans := runningSum([]int{1, 2, 3, 4})
	assert.Equal(t, []int{1, 3, 6, 10}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func runningSum1(nums []int) []int {
	ans := make([]int, len(nums)+1)
	for i, n := range nums {
		ans[i+1] = ans[i] + n
	}
	return ans[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
