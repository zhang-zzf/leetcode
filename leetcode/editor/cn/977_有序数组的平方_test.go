package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when977_thenSuccess(t *testing.T) {
	ans := sortedSquares([]int{-4, -1, 0, 3, 10})
	assert.Equal(t, []int{0, 1, 9, 16, 100}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
// 提前知道 ans 的长度，不需要反转 ans
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for l, r, idx := 0, n-1, n-1; l <= r; {
		ls, rs := nums[l]*nums[l], nums[r]*nums[r]
		if ls >= rs {
			ans[idx] = ls
			l, idx = l+1, idx-1
		} else {
			ans[idx] = rs
			r, idx = r-1, idx-1
		}
	}
	return ans
}

func sortedSquares1(nums []int) []int {
	var ans []int
	for l, r := 0, len(nums)-1; l <= r; {
		ls, rs := nums[l]*nums[l], nums[r]*nums[r]
		if ls >= rs {
			ans = append(ans, ls)
			l += 1
		} else {
			ans = append(ans, rs)
			r -= 1
		}
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
