package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1475_thenSuccess(t *testing.T) {
	ans := finalPrices([]int{8, 4, 6, 2, 3})
	assert.Equal(t, []int{4, 2, 4, 2, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func finalPrices(prices []int) []int {
	// TODO 0 做为哨兵，很关键
	rightMin := []int{0}
	for i := len(prices) - 1; i >= 0; i-- {
		price := prices[i]
		for len(rightMin) > 0 && price < rightMin[len(rightMin)-1] {
			rightMin = rightMin[:len(rightMin)-1]
		}
		prices[i] = price - rightMin[len(rightMin)-1]
		rightMin = append(rightMin, price)
	}
	return prices
}

//leetcode submit region end(Prohibit modification and deletion)
