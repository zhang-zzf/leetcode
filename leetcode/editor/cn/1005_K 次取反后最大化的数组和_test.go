package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1005_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, largestSumAfterKNegations([]int{4, 2, 3}, 1))
	assert.Equal(t, 6, largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	assert.Equal(t, 13, largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	assert.Equal(t, 15, largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 3))
	assert.Equal(t, 13, largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 4))
	assert.Equal(t, 15, largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 5))
}

func Test_givenFailedCase1_when1005_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, largestSumAfterKNegations([]int{-4, -2, -3}, 4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	sort.Ints(nums)
	for _, n := range nums {
		sum += n
	}
	i := 0
	for ; i < len(nums) && nums[i] < 0; i++ {
		if k <= 0 {
			break
		}
		sum += 2 * -nums[i]
		k -= 1
	}
	if k == 0 {
		return sum
	} else if k&0x01 == 1 {
		if i == len(nums) {
			sum -= 2 * -nums[i-1]
		} else {
			min := nums[i]
			if i-1 >= 0 && -nums[i-1] < nums[i] {
				min = -nums[i-1]
			}
			sum -= 2 * min
		}
	}
	return sum
}

// 算法是错误的 k > len(nums)
func largestSumAfterKNegationsOrigin(nums []int, k int) int {
	sum := 0
	sort.Ints(nums)
	for _, n := range nums {
		sum += n
	}
	for i := 0; i < k; i++ {
		if nums[i] < 0 {
			sum += -nums[i] * 2
		} else {
			if (k-i)&0x01 == 1 {
				sub := nums[i]
				if i-1 >= 0 && -nums[i-1] < nums[i] {
					sub = -nums[i-1]
				}
				sum -= sub * 2
			}
			break
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
