package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when867_thenSuccess(t *testing.T) {
	params := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	assert.Equal(t, [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}, transpose(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		column := make([]int, m)
		for j := 0; j < m; j++ {
			column[j] = matrix[j][i]
		}
		ans[i] = column
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
