package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when747_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, dominantIndex([]int{3, 6, 1, 0}))
	assert.Equal(t, -1, dominantIndex([]int{1, 2, 3, 4}))
	assert.Equal(t, 0, dominantIndex([]int{1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	m1, m2 := 0, 1
	if nums[m1] < nums[m2] {
		m1, m2 = m2, m1
	}
	for i := 2; i < len(nums); i++ {
		n := nums[i]
		if n > nums[m1] {
			m1, m2 = i, m1
		} else if n > nums[m2] {
			m2 = i
		}
	}
	if nums[m1] >= nums[m2]*2 {
		return m1
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
