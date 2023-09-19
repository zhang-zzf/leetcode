package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1295_thenSuccess(t *testing.T) {
	ans := findNumbers([]int{12, 345, 2, 6, 7896})
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findNumbers(nums []int) int {
	cntNum := func(a int) int {
		ans := 0
		for a > 0 {
			ans += 1
			a /= 10
		}
		return ans
	}
	ans := 0
	for _, n := range nums {
		if cntNum(n)&0x01 == 0 {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
