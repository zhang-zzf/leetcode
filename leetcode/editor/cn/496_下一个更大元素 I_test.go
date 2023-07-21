package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when496_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mapping := map[int]int{}
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		bigger := -1
		for lng := len(stack); lng > 0; lng = len(stack) {
			top := stack[lng-1]
			if top > n {
				bigger = top
				break
			}
			stack = stack[:lng-1]
		}
		stack = append(stack, n)
		mapping[n] = bigger
	}
	var ans []int
	for _, n := range nums1 {
		ans = append(ans, mapping[n])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
