package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when942_thenSuccess(t *testing.T) {
	assert.Equal(t, []int{0, 4, 1, 3, 2}, diStringMatch("IDID"))
	assert.Equal(t, []int{0, 1, 2, 3}, diStringMatch("III"))
	assert.Equal(t, []int{3, 2, 0, 1}, diStringMatch("DDI"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 若可以提前知道 []int{} 的长度，使用 make([]int,n) 比 append 效率更高
func diStringMatch(s string) []int {
	n := len(s)
	l, r := 0, n
	ans := make([]int, n+1)
	for i, c := range s {
		if c == 'I' {
			ans[i] = l
			l += 1
		} else {
			ans[i] = r
			r -= 1
		}
	}
	ans[n] = l
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
