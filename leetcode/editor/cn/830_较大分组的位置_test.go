package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when830_thenSuccess(t *testing.T) {
	ans := largeGroupPositions("abcdddeeeeaabbbcd")
	assert.Equal(t, [][]int{{3, 5}, {6, 9}, {12, 14}}, ans)
}

// 边界条件
// aaa
func Test_givenFailedCase1_when830_thenSuccess(t *testing.T) {
	ans := largeGroupPositions("aaa")
	assert.Equal(t, [][]int{{0, 2}}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func largeGroupPositions(s string) [][]int {
	var ans [][]int
	s += " "
	for l, r := 0, 1; r < len(s); r++ {
		if s[r] == s[r-1] {
			continue
		}
		if r-l > 2 {
			ans = append(ans, []int{l, r - 1})
		}
		l = r
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
