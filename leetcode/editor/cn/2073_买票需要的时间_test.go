package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2073_thenSuccess(t *testing.T) {
	assert.Equal(t, 6, timeRequiredToBuy([]int{2, 3, 2}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func timeRequiredToBuy(tickets []int, k int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	ans := 0
	pivot := tickets[k]
	for i, n := range tickets {
		if i <= k {
			ans += min(pivot, n)
		} else {
			ans += min(pivot-1, n)
		}
	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
