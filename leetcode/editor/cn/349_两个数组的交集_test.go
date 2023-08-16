package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when349_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	set := map[int]int{}
	for _, num := range nums1 {
		set[num] = 1
	}
	// 思路：遍历 nums2 时更新 ans
	for _, num := range nums2 {
		// TODO set[notExistsKey] = 0
		if set[num] > 0 {
			ans = append(ans, num)
			set[num] = 0
		}
	}
	return ans
}

func intersection1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	set := map[int]int{}
	for _, num := range nums1 {
		set[num] = 1
	}
	for _, num := range nums2 {
		if cnt, ok := set[num]; ok {
			set[num] = cnt + 1
		}
	}
	var ans []int
	for k, v := range set {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
