package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2078_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, maxDistance([]int{1, 1, 1, 6, 1, 1, 1}))
	assert.Equal(t, 8, maxDistance([]int{4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(colors []int) int {
	ans := 0
	n := len(colors)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if colors[i] != colors[j] {
				if j-i > ans {
					ans = j - i
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
