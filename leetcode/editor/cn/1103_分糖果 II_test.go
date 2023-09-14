package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1103_thenSuccess(t *testing.T) {
	ans := distributeCandies(10, 3)
	assert.Equal(t, []int{5, 2, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	for idx := 0; candies > 0; idx++ {
		cnt := idx + 1
		if cnt > candies {
			cnt = candies
		}
		candies -= cnt
		ans[idx%num_people] += cnt
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
