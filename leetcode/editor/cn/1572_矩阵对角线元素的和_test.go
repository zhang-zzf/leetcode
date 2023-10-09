package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1572_thenSuccess(t *testing.T) {
	params := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	assert.Equal(t, 25, diagonalSum(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func diagonalSum(mat [][]int) int {
	m := len(mat)
	ans := 0
	for i := 0; i < m; i++ {
		ans += mat[i][i] + mat[i][m-1-i]
	}
	if m&0x01 == 0x01 {
		ans -= mat[m/2][m/2]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
