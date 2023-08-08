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
	for i, item := range nums {
		for j := 1; j <= k; j++ {
			if i+j < len(nums) && item == nums[i+j] {
				ans = true
				break
			}
		}
		if ans == true {
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
