package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1403_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	ts, ls := 0, 0
	for _, n := range nums {
		ts += n
	}
	for idx, n := range nums {
		ls += n
		if ls > ts/2 {
			return nums[:idx+1]
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
