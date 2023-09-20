package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1313_thenSuccess(t *testing.T) {
	ans := decompressRLElist([]int{1, 2, 3, 4})
	assert.Equal(t, []int{2, 4, 4, 4}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func decompressRLElist(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
