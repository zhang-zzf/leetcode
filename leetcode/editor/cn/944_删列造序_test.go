package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when944_thenSuccess(t *testing.T) {
	ans := minDeletionSize([]string{"cba", "daf", "ghi"})
	assert.Equal(t, 1, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minDeletionSize(strs []string) int {
	ans := 0
	m, n := len(strs), len(strs[0])
	for c := 0; c < n; c++ {
		for r := 0; r+1 < m; r++ {
			if strs[r][c] > strs[r+1][c] {
				ans += 1
				break
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
