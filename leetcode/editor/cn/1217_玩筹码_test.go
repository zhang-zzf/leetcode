package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1217_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, minCostToMoveChips([]int{2, 2, 2, 3, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minCostToMoveChips(position []int) int {
	// 把所有的都移动到1或者都移动到2，取2者的较小值
	n1, n2 := 0, 0
	for _, n := range position {
		if n&0x01 == 0 {
			n1 += 1
		} else {
			n2 += 1
		}
	}
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	return n1
}

//leetcode submit region end(Prohibit modification and deletion)
