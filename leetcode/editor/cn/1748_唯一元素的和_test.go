package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1748_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumOfUnique(nums []int) int {
	cnt := make([]int, 101)
	for _, n := range nums {
		cnt[n] += 1
	}
	ans := 0
	for n, c := range cnt {
		if c == 1 {
			ans += n
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
