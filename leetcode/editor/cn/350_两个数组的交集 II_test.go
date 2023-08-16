package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when350_thenSuccess(t *testing.T) {
	ans := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	assert.Equal(t, []int{2, 2}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	var cntMap = map[int]int{}
	for _, x := range nums1 {
		cntMap[x] += 1
	}
	for _, x := range nums2 {
		if cntMap[x] > 0 {
			ans = append(ans, x)
			cntMap[x] -= 1
		}
	}
	return ans
}

func intersect2(nums1 []int, nums2 []int) []int {
	var ans []int
	// TODO sort
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i += 1
			continue
		} else if nums1[i] > nums2[j] {
			j++
			continue
		}
		ans = append(ans, nums1[i])
		i, j = i+1, j+1
	}
	return ans
}

func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// TODO go 值传递
	// 值类型 - int、float、bool、string、array、struct等
	// 引用类型 - slice，map，channel，interface，func等
	cntMap := map[int]*[2]int{}
	for _, x := range nums1 {
		if cnt, ok := cntMap[x]; ok {
			cnt[0] += 1
		} else {
			cntMap[x] = &[2]int{1, 0}
		}
	}
	for _, x := range nums2 {
		if cnt, ok := cntMap[x]; ok {
			cnt[1] += 1
		}
	}
	var ans []int
	for k, v := range cntMap {
		cnt := v[0]
		if cnt > v[1] {
			cnt = v[1]
		}
		for i := 0; i < cnt; i++ {
			ans = append(ans, k)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
