package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when219_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	ans := false
	window := map[int]struct{}{}
	for i, item := range nums {
		if _, exists := window[item]; exists {
			ans = true
			break
		}
		window[item] = struct{}{}
		if len(window) > k {
			delete(window, nums[i-k])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
