package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when896_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isMonotonic([]int{1, 2, 2, 3}))
	assert.Equal(t, true, isMonotonic([]int{3, 2}))
	assert.Equal(t, false, isMonotonic([]int{3, 1, 2}))
	assert.Equal(t, true, isMonotonic([]int{2, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isMonotonic(nums []int) bool {
	incr, decr := true, true
	for i := 1; i < len(nums); i++ {
		m := nums[i] - nums[i-1]
		if m > 0 {
			decr = false
		} else if m < 0 {
			incr = false
		}
		if !(incr || decr) {
			break
		}
	}
	return incr || decr
}

func isMonotonic1(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	monotonic := 0
	ans := true
	for i := 1; i < n; i++ {
		m := nums[i] - nums[i-1]
		if monotonic != 0 {
			if monotonic*m < 0 {
				ans = false
			}
			continue
		}
		switch {
		case m > 0:
			monotonic = 1
		case m < 0:
			monotonic = -1
		default:
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
