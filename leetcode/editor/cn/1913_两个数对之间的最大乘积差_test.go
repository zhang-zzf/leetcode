package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1913_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}

//leetcode submit region end(Prohibit modification and deletion)
