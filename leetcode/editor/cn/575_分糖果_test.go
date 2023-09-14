package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when575_thenSuccess(t *testing.T) {
	ans := distributeCandies575([]int{1, 1, 2, 3})
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies575(candyType []int) int {
	diffMap := map[int]struct{}{}
	for _, t := range candyType {
		diffMap[t] = struct{}{}
	}
	ans := len(diffMap)
	if ans > len(candyType)/2 {
		ans = len(candyType) / 2
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
