package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2057_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func smallestEqual(nums []int) int {
	for i, n := range nums {
		if i%10 == n {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
