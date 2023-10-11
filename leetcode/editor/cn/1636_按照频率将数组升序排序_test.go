package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1636_thenSuccess(t *testing.T) {
	ans := frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1})
	assert.Equal(t, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(nums []int) []int {
	numToCnt := make([]int, 210)
	for _, num := range nums {
		numToCnt[num+100] += 1
	}
	sort.Slice(nums, func(i, j int) bool {
		compare := numToCnt[nums[i]+100] - numToCnt[nums[j]+100]
		if compare < 0 {
			return true
		} else if compare == 0 {
			return nums[i] > nums[j]
		}
		return false
	})
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
