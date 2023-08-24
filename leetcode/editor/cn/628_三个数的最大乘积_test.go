package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when628_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumProduct(nums []int) int {
	// 思路 3个最大的数 或者 最大的数*2个最小的数
	sort.Ints(nums)
	last := len(nums) - 1
	m1 := nums[0] * nums[1] * nums[last]
	m2 := nums[last] * nums[last-1] * nums[last-2]
	if m1 > m2 {
		return m1
	} else {
		return m2
	}
}

//leetcode submit region end(Prohibit modification and deletion)
