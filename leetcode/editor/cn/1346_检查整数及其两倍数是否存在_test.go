package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1346_thenSuccess(t *testing.T) {
	assert.Equal(t, true, checkIfExist([]int{10, 1, 5, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkIfExist(arr []int) bool {
	set := map[int]int{}
	for i, n := range arr {
		set[n] = i
	}
	for i, n := range arr {
		if idx, ok := set[n*2]; ok && i != idx {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
