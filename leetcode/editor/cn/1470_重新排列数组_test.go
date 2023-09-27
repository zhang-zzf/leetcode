package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1470_thenSuccess(t *testing.T) {
	ans := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	assert.Equal(t, []int{2, 3, 5, 4, 1, 7}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shuffle(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		if i&0x01 == 0 {
			ans[i] = nums[i/2]
		} else {
			ans[i] = nums[n+i/2]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
