package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when796_thenSuccess(t *testing.T) {
	assert.Equal(t, true, rotateString("abcde", "cdeab"))
	assert.Equal(t, false, rotateString("abcde", "cedab"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 拼接2个s
func rotateString1(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

func rotateString(s string, goal string) bool {
	// 模拟旋转
	m, n := len(s), len(goal)
	if m != n {
		return false
	}
	ans := false
	for i, j := 0, 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if s[(i+j)%n] != goal[j] {
				break
			}
		}
		if j == n {
			ans = true
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
