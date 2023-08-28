package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when717_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isOneBitCharacter([]int{1, 0, 0}))
	assert.Equal(t, false, isOneBitCharacter([]int{1, 1, 1, 0}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isOneBitCharacter(bits []int) bool {
	ans := false
	n := len(bits)
	for i := 0; i < n; i++ {
		if 1 == bits[i] {
			i++
		} else if i == n-1 {
			ans = true
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
