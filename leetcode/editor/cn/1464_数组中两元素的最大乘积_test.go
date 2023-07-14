package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1464_thenSuccess(t *testing.T) {
	assert.Equal(t, 12, maxProduct([]int{3, 4, 5, 2}), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	m1, m2 := nums[0], nums[1]
	if m1 < m2 {
		m1, m2 = m2, m1
	}
	for _, n := range nums[2:] {
		if n > m1 {
			m1, m2 = n, m1
		} else if n > m2 {
			m2 = n
		}
	}
	return (m1 - 1) * (m2 - 1)
}

//leetcode submit region end(Prohibit modification and deletion)
