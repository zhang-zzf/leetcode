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
	// map used as a Set
	window := map[int]struct{}{}
	for i, item := range nums {
		// check item in Set
		if _, exists := window[item]; exists {
			ans = true
			break
		}
		// add to Set
		window[item] = struct{}{}
		if len(window) > k {
			// remove from Set
			delete(window, nums[i-k])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
