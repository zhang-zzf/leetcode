package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1037_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isBoomerang(points [][]int) bool {
	// 几何向量
	v1 := []int{points[1][0] - points[0][0], points[1][1] - points[0][1]}
	v2 := []int{points[2][0] - points[0][0], points[2][1] - points[0][1]}
	return v1[0]*v2[1]-v1[1]*v2[0] != 0
}

//leetcode submit region end(Prohibit modification and deletion)
