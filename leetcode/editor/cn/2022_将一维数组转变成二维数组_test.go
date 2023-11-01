package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2022_thenSuccess(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}}, construct2DArray([]int{1, 2, 3}, 1, 3))
}

func Test_givenFailedCase1_when2022_thenSuccess(t *testing.T) {
	assert.Equal(t, [][]int{{5}}, construct2DArray([]int{5}, 3, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func construct2DArray(original []int, m int, n int) [][]int {
	var ans [][]int
	if m*n != len(original) {
		return ans
	}
	var cur []int
	for i := 0; i < len(original); i++ {
		cur = append(cur, original[i])
		if (i+1)%n == 0 {
			ans = append(ans, cur)
			cur = []int{}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
