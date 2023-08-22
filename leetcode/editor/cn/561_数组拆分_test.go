package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when561_thenSuccess(t *testing.T) {
	ans := arrayPairSum([]int{6, 2, 6, 5, 1, 2})
	assert.Equal(t, 9, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func arrayPairSum(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
