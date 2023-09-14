package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1089_thenSuccess(t *testing.T) {
	param := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(param)
	assert.Equal(t, []int{1, 0, 0, 2, 3, 0, 0, 4}, param)
}

func Test_givenFailedCase1_when1089_thenSuccess(t *testing.T) {
	param := []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(param)
	assert.Equal(t, []int{8, 4, 5, 0, 0, 0, 0, 0}, param)
}

//leetcode submit region begin(Prohibit modification and deletion)
func duplicateZeros(arr []int) {
	m := 0
	for _, v := range arr {
		if v == 0 {
			m += 1
		}
	}
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 0 {
			m -= 1
		}
		if i+m < n {
			arr[i+m] = arr[i]
			if arr[i] == 0 && i+m+1 < n {
				arr[i+m+1] = 0
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
