package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when2089_thenSuccess(t *testing.T) {
	ans := targetIndices([]int{1, 2, 5, 2, 3}, 5)
	assert.Equal(t, []int{4}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ans := make([]int, 0)
	for i, n := range nums {
		if n == target {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
