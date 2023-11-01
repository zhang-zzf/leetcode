package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func Test_givenNormal_when1984_thenSuccess(t *testing.T) {
	ans := minimumDifference([]int{9, 4, 1, 7}, 2)
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minimumDifference(nums []int, k int) int {
	if k < 2 {
		return 0
	}
	// sort then slide window
	sort.Ints(nums)
	ans := math.MaxInt
	for i := k - 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-k+1]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
