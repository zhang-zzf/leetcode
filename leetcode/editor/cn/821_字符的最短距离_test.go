package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when821_thenSuccess(t *testing.T) {
	ans := shortestToChar("loveleetcode", 'e')
	assert.Equal(t, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = math.MaxInt
	}
	last := -1
	for i := 0; i < n; i++ {
		if c == s[i] {
			last = i
		}
		period := i - last
		if last != -1 && period < ans[i] {
			ans[i] = period
		}
	}
	last = -1
	for i := n - 1; i >= 0; i-- {
		if c == s[i] {
			last = i
		}
		period := last - i
		if last != -1 && period < ans[i] {
			ans[i] = period
		}
	}
	return ans
}

func shortestToChar1(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = math.MaxInt32
	}
	for i, char := range s {
		if c == byte(char) {
			ans[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		if ans[i] != 0 {
			continue
		}
		// 向左
		for j := 1; i-j >= 0 && ans[i-j] > j; j++ {
			ans[i-j] = j
		}
		// 向右
		for j := 1; i+j < n && ans[i+j] > j; j++ {
			ans[i+j] = j
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
