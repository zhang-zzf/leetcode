package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1287_thenSuccess(t *testing.T) {
	assert.Equal(t, 6, findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}

func Test_givenFailedCase1_when1287_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, findSpecialInteger([]int{1}))
	assert.Equal(t, 1, findSpecialInteger([]int{1, 1}))
	assert.Equal(t, 3, findSpecialInteger([]int{1, 2, 3, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findSpecialInteger(arr []int) (ans int) {
	n := len(arr)
	if n < 4 {
		return arr[0]
	}
	p25, nn := n/4, 1
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			nn = 1
		} else {
			nn += 1
		}
		if nn > p25 {
			ans = arr[i]
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
