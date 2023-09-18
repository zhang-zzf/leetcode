package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1232_thenSuccess(t *testing.T) {
	params := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	assert.Equal(t, true, checkStraightLine(params))
}

func Test_givenFailedCase1_when1232_thenSuccess(t *testing.T) {
	params := [][]int{{1, 2}, {2, 3}, {3, 5}}
	assert.Equal(t, false, checkStraightLine(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkStraightLine(c [][]int) bool {
	// 向量
	n := len(c)
	v1n := []int{c[n-1][0] - c[0][0], c[n-1][1] - c[0][1]}
	for i := 1; i < n; i++ {
		v1i := []int{c[i][0] - c[0][0], c[i][1] - c[0][1]}
		if v1n[0]*v1i[1]-v1n[1]*v1i[0] != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
