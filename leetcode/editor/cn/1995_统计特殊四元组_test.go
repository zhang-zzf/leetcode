package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1995_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, countQuadruplets([]int{1, 1, 1, 3, 5}))
	assert.Equal(t, 0, countQuadruplets([]int{3, 3, 6, 4, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countQuadruplets(nums []int) int {
	// 爆破
	ans := 0
	n := len(nums)
	for i := 3; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := j + 1; k < i; k++ {
				for l := k + 1; l < i; l++ {
					if nums[i] == nums[j]+nums[k]+nums[l] {
						ans += 1
					}
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
