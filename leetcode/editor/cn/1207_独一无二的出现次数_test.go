package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1207_thenSuccess(t *testing.T) {
	assert.Equal(t, true, uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func uniqueOccurrences(arr []int) bool {
	numCnt := map[int]int{}
	for _, n := range arr {
		numCnt[n] += 1
	}
	cntSet := map[int]int{}
	for _, cnt := range numCnt {
		if cntSet[cnt] += 1; cntSet[cnt] > 1 {
			return false
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)
