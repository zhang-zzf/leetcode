package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1394_thenSuccess(t *testing.T) {
	ans := findLucky([]int{1, 2, 2, 3, 3, 3})
	assert.Equal(t, 3, ans)
}

func Test_givenFailedCase1_when1394_thenSuccess(t *testing.T) {
	ans := findLucky([]int{2, 2, 2, 3, 3})
	assert.Equal(t, -1, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findLucky(arr []int) int {
	mapping := make([]int, 501)
	for _, n := range arr {
		mapping[n] += 1
	}
	for i := 500; i >= 0; i-- {
		if i != 0 && i == mapping[i] {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
