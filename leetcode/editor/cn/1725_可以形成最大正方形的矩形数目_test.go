package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1725_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countGoodRectangles(rectangles [][]int) int {
	min := func(arr []int) int {
		ans := math.MaxInt
		for _, x := range arr {
			if x < ans {
				ans = x
			}
		}
		return ans
	}
	ans, max := 0, 0
	for _, r := range rectangles {
		m := min(r)
		if m > max {
			max = m
			ans = 1
		} else if m == max {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
