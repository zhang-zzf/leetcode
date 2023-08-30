package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when766_thenSuccess(t *testing.T) {
	params := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	assert.Equal(t, true, isToeplitzMatrix(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

// 想复杂了
func isToeplitzMatrix1(matrix [][]int) bool {
	checkMatrix := func(matrix [][]int, y int, x int, n int, m int) bool {
		pivot := matrix[y][x]
		for ; x < n && y < m; x, y = x+1, y+1 {
			if matrix[y][x] != pivot {
				return false
			}
		}
		return true
	}
	m, n := len(matrix), len(matrix[0])
	// row
	for i := 0; i < m; i++ {
		if checkMatrix(matrix, i, 0, n, m) == false {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if checkMatrix(matrix, 0, i, n, m) == false {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
