package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1848_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func getMinDistance(nums []int, target int, start int) int {
	ans := 0
	n := len(nums)
	for ans < n {
		l, r := start-ans, start+ans
		if l >= 0 && nums[l] == target {
			break
		}
		if r < n && nums[r] == target {
			break
		}
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
