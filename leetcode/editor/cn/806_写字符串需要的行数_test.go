package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when806_thenSuccess(t *testing.T) {
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	ans := numberOfLines(widths, "bbbcccdddaaa")
	assert.Equal(t, []int{2, 4}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfLines(widths []int, s string) []int {
	ans := []int{1, 0}
	for _, c := range s {
		w := widths[c-'a']
		ans[1] += w
		if ans[1] > 100 {
			ans[0] += 1
			// error
			// ans[1] -= 100
			ans[1] = w
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
