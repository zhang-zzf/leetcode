package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when922_thenSuccess(t *testing.T) {
	ans := sortArrayByParityII([]int{4, 2, 5, 7})
	assert.Equal(t, []int{4, 5, 2, 7}, ans)
}

func Test_givenFailedCase1_when922_thenSuccess(t *testing.T) {
	ans := sortArrayByParityII([]int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3})
	assert.Equal(t, []int{2, 3, 0, 1, 4, 1, 0, 3, 4, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(nums []int) []int {
	e, o, n := 0, 1, len(nums)
	for e < n && o < n {
		for e < n && nums[e]&0x01 == 0 {
			e += 2
		}
		for o < n && nums[o]&0x01 == 1 {
			o += 2
		}
		if e < n && o < n {
			nums[e], nums[o] = nums[o], nums[e]
		}
	}
	return nums
}

// 算法是错误的
func sortArrayByParityII1(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		if (l%2)^(nums[l]%2) == 0 {
			l += 1
			continue
		}
		if (r%2)^(nums[r]%2) == 0 {
			r -= 1
			continue
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
