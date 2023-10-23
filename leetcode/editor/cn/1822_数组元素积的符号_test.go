package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1822_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	assert.Equal(t, 0, arraySign([]int{-0, -2, -3, -4, 3, 2, 1}))
	assert.Equal(t, -1, arraySign([]int{-1, 2, -3, -4, 3, 2, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func arraySign(nums []int) int {
	negative := 0
	for _, n := range nums {
		if n < 0 {
			negative += 1
		} else if n == 0 {
			return 0
		}
	}
	if negative&0x01 == 0 {
		return 1
	} else {
		return -1
	}

}

//leetcode submit region end(Prohibit modification and deletion)
