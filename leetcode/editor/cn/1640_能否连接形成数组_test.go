package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1640_thenSuccess(t *testing.T) {
	ans := canFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}})
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func canFormArray(arr []int, pieces [][]int) bool {
	var piece []int
	ptr := 0
	for _, num := range arr {
		if ptr == len(piece) {
			// re find the piece
			for _, p := range pieces {
				if p[0] == num {
					piece = p
					ptr = 0
					break
				}
			}
		}
		if ptr >= len(piece) || num != piece[ptr] {
			return false
		}
		ptr += 1
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
