package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when_thenSuccess(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0}
	merge(nums, 3, []int{2, 5, 6}, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for idx := len(nums1) - 1; idx >= 0; idx-- {
		if n-1 < 0 {
			break
		}
		if m-1 >= 0 && nums1[m-1] >= nums2[n-1] {
			nums1[idx] = nums1[m-1]
			m -= 1
		} else {
			nums1[idx] = nums2[n-1]
			n -= 1
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
