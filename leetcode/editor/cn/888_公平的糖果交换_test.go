package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when888_thenSuccess(t *testing.T) {
	ans := fairCandySwap([]int{1, 2, 5}, []int{2, 4})
	assert.Equal(t, []int{5, 4}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	// alice give x, bob give y
	m, n := 0, 0
	bobCnt := map[int]struct{}{}
	for _, i := range aliceSizes {
		m += i
	}
	for _, i := range bobSizes {
		n += i
		bobCnt[i] = struct{}{}
	}
	// y = x + (n-m)/2
	subtract := (n - m) / 2
	ans := make([]int, 2)
	for _, x := range aliceSizes {
		y := x + subtract
		if _, ok := bobCnt[y]; ok {
			ans[0], ans[1] = x, y
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
