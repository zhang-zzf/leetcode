package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1380_thenSuccess(t *testing.T) {
	ans := luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}})
	assert.Equal(t, []int{15}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func luckyNumbers(matrix [][]int) []int {
	var rowMin []int
	for _, row := range matrix {
		min := row[0]
		for _, n := range row {
			if n < min {
				min = n
			}
		}
		rowMin = append(rowMin, min)
	}
	var ans []int
	for c := 0; c < len(matrix[0]); c++ {
		maxIdx := 0
		for r := 0; r < len(matrix); r++ {
			if matrix[r][c] > matrix[maxIdx][c] {
				maxIdx = r
			}
		}
		if matrix[maxIdx][c] == rowMin[maxIdx] {
			ans = append(ans, matrix[maxIdx][c])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
