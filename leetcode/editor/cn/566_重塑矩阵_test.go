package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when566_thenSuccess(t *testing.T) {
	params := [][]int{
		{1, 2},
		{3, 4},
	}
	ans := matrixReshape(params, 1, 4)
	assert.Equal(t, [][]int{{1, 2, 3, 4}}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	ans := make([][]int, r)
	rc, cc := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cc == c {
				cc = 0
				rc += 1
			}
			if cc == 0 {
				ans[rc] = make([]int, c)
			}
			ans[rc][cc] = mat[i][j]
			cc += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
