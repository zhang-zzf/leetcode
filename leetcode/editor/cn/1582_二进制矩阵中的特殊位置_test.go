package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1582_thenSuccess(t *testing.T) {
	params := [][]int{
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0}}
	assert.Equal(t, 2, numSpecial1582(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 没有什么特出算法，最简单朴实的模拟就是最好的
func numSpecial1582(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rowCnt := make([]int, m)
	for r, row := range mat {
		cnt := 0
		for _, num := range row {
			if num == 1 {
				cnt += 1
			}
		}
		rowCnt[r] = cnt
	}
	ans := 0
	for c := 0; c < n; c++ {
		oneCnt, row := 0, -1
		for r := 0; r < m; r++ {
			if mat[r][c] == 1 {
				oneCnt, row = oneCnt+1, r
			}
		}
		if oneCnt == 1 && rowCnt[row] == 1 {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
