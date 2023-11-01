package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2006_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, countKDifference([]int{3, 2, 1, 5, 4}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countKDifference(nums []int, k int) int {
	// hash 以数组表示
	ans := 0
	cnt := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		if m := nums[i] - k; m >= 0 {
			ans += cnt[m]
		}
		if m := nums[i] + k; m < 101 {
			ans += cnt[m]
		}
		cnt[nums[i]] += 1
	}
	return ans
}

func countKDifference2(nums []int, k int) int {
	// hash
	ans := 0
	cnt := map[int]int{}
	for i := 0; i < len(nums); i++ {
		ans += cnt[nums[i]-k] + cnt[nums[i]+k]
		cnt[nums[i]] += 1
	}
	return ans
}

func countKDifference1(nums []int, k int) int {
	// 爆破
	ans := 0
	for r := 1; r < len(nums); r++ {
		for l := 0; l < r; l++ {
			if nums[l]+k == nums[r] || nums[l]-k == nums[r] {
				ans += 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
