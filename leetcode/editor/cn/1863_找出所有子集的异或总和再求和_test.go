package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1863_thenSuccess(t *testing.T) {
	assert.Equal(t, 28, subsetXORSum([]int{5, 1, 6}))
	assert.Equal(t, 480, subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 回溯算法
func subsetXORSum(nums []int) int {
	n := len(nums)
	sum := 0
	var backTrace func(int, int)
	backTrace = func(index int, xor int) {
		for i := index; i < n; i++ {
			// todo 画树图
			cxor := xor ^ nums[i]
			sum += cxor
			backTrace(i+1, cxor)
		}
	}
	backTrace(0, 0)
	return sum
}

// 算法是错误的
// 3，4，5，6 选择不到 [3,4,6]
func subsetXORSumErr(nums []int) int {
	// 分而治之
	n := len(nums)
	if n <= 0 {
		return 0
	}
	sum := nums[0]
	for i := 1; i < n; i++ {
		xor := nums[0]
		for j := i; j < n; j++ {
			xor ^= nums[j]
			sum += xor
		}
	}
	return sum + subsetXORSum(nums[1:])
}

//leetcode submit region end(Prohibit modification and deletion)
