package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1431_thenSuccess(t *testing.T) {
	ans := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	assert.Equal(t, []bool{true, true, true, false, true}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := math.MinInt
	for _, n := range candies {
		if n > max {
			max = n
		}
	}
	num := max - extraCandies
	ans := make([]bool, len(candies))
	for i, n := range candies {
		ans[i] = n >= num
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
