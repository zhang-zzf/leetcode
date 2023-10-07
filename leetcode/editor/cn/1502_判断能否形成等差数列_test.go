package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1502_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canMakeArithmeticProgression([]int{1, 5, 3}))
	assert.Equal(t, false, canMakeArithmeticProgression([]int{1, 4, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	subtract := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != subtract {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
