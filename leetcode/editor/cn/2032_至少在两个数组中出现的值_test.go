package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2032_thenSuccess(t *testing.T) {
	ans := twoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5})
	assert.Equal(t, []int{}, ans)
}

func Test_givenFailedCase1_when2032_thenSuccess(t *testing.T) {
	ans := twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3})
	assert.Equal(t, []int{2, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	// todo 参考答案 使用 bit 统计
	cnt := make([]int, 101)
	for idx, arr := range [][]int{nums1, nums2, nums3} {
		for _, n := range arr {
			cnt[n] |= 1 << idx
		}
	}
	ans := make([]int, 0)
	//countOne := func(x int) int {
	//	ans := 0
	//	for x > 0 {
	//		ans += 1
	//		x &= x - 1
	//	}
	//	return ans
	//}
	//for n, c := range cnt {
	//	if countOne(c) > 1 {
	//		ans = append(ans, n)
	//	}
	//}
	for n, c := range cnt {
		if c&c-1 > 0 {
			ans = append(ans, n)
		}
	}
	return ans
}

func twoOutOfThree2(nums1 []int, nums2 []int, nums3 []int) []int {
	// hash 统计
	cnt := make([]int, 101)
	for _, arr := range [][]int{nums1, nums2, nums3} {
		set := make([]int, 101)
		for _, n := range arr {
			set[n] = 1
		}
		for n, c := range set {
			if c > 0 {
				cnt[n] += 1
			}
		}
	}
	ans := make([]int, 0)
	for n, c := range cnt {
		if c > 1 {
			ans = append(ans, n)
		}
	}
	return ans
}

func twoOutOfThree1(nums1 []int, nums2 []int, nums3 []int) []int {
	// hash 统计
	cnt := make([]int, 101)
	for _, n := range nums1 {
		cnt[n] = 1
	}
	// todo 算法是错误的
	for _, n := range nums2 {
		if cnt[n] == 0 {
			cnt[n] = 1
		} else if cnt[n] == 1 {
			cnt[n] = 2
		}
	}
	for _, n := range nums3 {
		if cnt[n] == 1 {
			cnt[n] = 2
		}
	}
	var ans []int
	for n, c := range cnt {
		if c == 2 {
			ans = append(ans, n)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
