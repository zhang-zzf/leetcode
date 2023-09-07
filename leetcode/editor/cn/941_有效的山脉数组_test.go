package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when941_thenSuccess(t *testing.T) {
	assert.Equal(t, true, validMountainArray([]int{0, 2, 3, 4, 5, 2, 0}))
	assert.Equal(t, false, validMountainArray([]int{2, 1}))
	assert.Equal(t, false, validMountainArray([]int{2, 5, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO
func validMountainArray(arr []int) bool {
	i, n := 0, len(arr)
	for ; i+1 < n && arr[i] < arr[i+1]; i++ {
	}
	if i == 0 || i == n-1 {
		return false
	}
	for ; i+1 < n; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}

// 最初最大值
func validMountainArray1(arr []int) bool {
	n := len(arr)
	max := 0
	for i := 0; i < n; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	if max == 0 || max == n-1 {
		return false
	}
	for i := 0; i < max; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	for i := n - 1; i > max; i-- {
		if arr[i] >= arr[i-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
