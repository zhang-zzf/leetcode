package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1566_thenSuccess(t *testing.T) {
	assert.Equal(t, true, containsPattern([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2))
	assert.Equal(t, false, containsPattern([]int{1, 2, 3, 1, 2}, 2, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func containsPattern(arr []int, m int, k int) bool {
	n := len(arr)
	for i := 0; i <= n-m*k; i++ {
		j := 0
		for ; j < m*k; j++ {
			if arr[i+j] != arr[i+j%m] {
				break
			}
		}
		if j == m*k {
			return true
		}
	}
	return false
}

func containsPattern1(arr []int, m int, k int) bool {
	for i := 0; i < len(arr); i++ {
		j := 0
		for ; j < m; j++ {
			l := 1
			for ; l < k; l++ {
				idx := i + j + l*m
				if i+j >= len(arr) || idx >= len(arr) || arr[idx] != arr[i+j] {
					break
				}
			}
			if l != k {
				break
			}
		}
		if j == m {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
