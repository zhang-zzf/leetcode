package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1886_thenSuccess(t *testing.T) {
	mat := [][]int{{0, 1}, {1, 0}}
	target := [][]int{{1, 0}, {0, 1}}
	assert.Equal(t, true, findRotation(mat, target))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	matrixrc := func(m [][]int, r, c, rc int) int {
		switch rc {
		case 0:
			return m[r][c]
		case 1:
			return m[n-1-c][r]
		case 2:
			return m[n-1-r][n-1-c]
		case 3:
			return m[c][n-1-r]
		default:
			return m[r][c]
		}
	}
	same := func(m [][]int, rc int) bool {
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if matrixrc(mat, r, c, rc) != target[r][c] {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < 4; i++ {
		if same(mat, i) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
