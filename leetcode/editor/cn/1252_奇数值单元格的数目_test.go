package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1252_thenSuccess(t *testing.T) {
	ans := oddCells(2, 2, [][]int{{1, 1}, {0, 0}})
	assert.Equal(t, 0, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func oddCells(m int, n int, indices [][]int) (ans int) {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for _, rc := range indices {
		r, c := rc[0], rc[1]
		for i := 0; i < n; i++ {
			dp[r][i] += 1
		}
		for i := 0; i < m; i++ {
			dp[i][c] += 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j]&0x01 == 1 {
				ans += 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
