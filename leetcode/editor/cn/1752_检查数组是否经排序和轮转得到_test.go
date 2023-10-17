package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1752_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func check(nums []int) bool {
	reversed, n := 0, len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			reversed += 1
		}
	}
	if nums[0] < nums[n-1] {
		return reversed == 0
	}
	return reversed <= 1
}

//leetcode submit region end(Prohibit modification and deletion)
