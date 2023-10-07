package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1512_thenSuccess(t *testing.T) {
	ans := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	assert.Equal(t, 4, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力破
func numIdenticalPairs(nums []int) int {
	ans := 0
	mapping := map[int]int{}
	for _, n := range nums {
		ans += mapping[n]
		mapping[n] += 1
	}
	return ans
}

func numIdenticalPairs1(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				ans += 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
