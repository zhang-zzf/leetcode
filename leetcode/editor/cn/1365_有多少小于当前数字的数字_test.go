package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1365_thenSuccess(t *testing.T) {
	ans := smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})
	assert.Equal(t, []int{4, 0, 1, 1, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func smallerNumbersThanCurrent11(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Ints(arr)
	digitCnt := map[int]int{}
	for i, n := range arr {
		if _, ok := digitCnt[n]; !ok {
			digitCnt[n] = i
		}
	}
	for i := 0; i < n; i++ {
		nums[i] = digitCnt[nums[i]]
	}
	return nums
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for _, num := range nums {
		for j := 0; j < n; j++ {
			if num < nums[j] {
				ans[j] += 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
