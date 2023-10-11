package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1652_thenSuccess(t *testing.T) {
	assert.Equal(t, []int{12, 10, 16, 13}, decrypt([]int{5, 7, 1, 4}, 3))
	assert.Equal(t, []int{12, 10, 16, 13}, decrypt([]int{5, 7, 1, 4}, -3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func decrypt(code []int, k int) []int {
	n := len(code)
	if k == 0 {
		return make([]int, n)
	}
	prefixSum := make([]int, n*2+1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + code[(i-1)%n]
	}
	for i := 0; i < n; i++ {
		l, r := i+1, i+1+k
		if k < 0 {
			l, r = i+n+k, i+n
		}
		code[i] = prefixSum[r] - prefixSum[l]
	}
	return code
}

//leetcode submit region end(Prohibit modification and deletion)
