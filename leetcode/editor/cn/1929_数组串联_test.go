package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1929_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func getConcatenation(nums []int) []int {
	var ans []int
	ans = append(ans, nums...)
	ans = append(ans, nums...)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
