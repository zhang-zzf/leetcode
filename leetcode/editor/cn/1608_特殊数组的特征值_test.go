package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1608_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, specialArray([]int{0, 4, 3, 0, 4}))
	assert.Equal(t, -1, specialArray([]int{3, 6, 7, 7, 0}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func specialArray(nums []int) int {
	// desc sort
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	n := len(nums)
	for i := 1; i <= n; i++ {
		if nums[i-1] < i {
			break
		}
		if i == n || nums[i] < i {
			return i
		}
	}
	return -1
}

func specialArray1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i <= n; i++ {
		cnt := n - sort.Search(n, func(j int) bool { return nums[j] >= i })
		if cnt == i {
			return i
		}
		if cnt < i {
			break
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
