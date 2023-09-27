package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1460_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canBeEqual([]int{1, 2, 3, 4}, []int{2, 3, 1, 4}))
	assert.Equal(t, false, canBeEqual([]int{1, 2, 3, 4, 4}, []int{2, 3, 1, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canBeEqual(target []int, arr []int) bool {
	// 数组基数统计更快
	numCnt := make([]int, 1001)
	for _, n := range target {
		numCnt[n] += 1
	}
	for _, n := range arr {
		if numCnt[n] == 0 {
			return false
		}
		numCnt[n] -= 1
	}
	for _, cnt := range numCnt {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func canBeEqual1(target []int, arr []int) bool {
	//  数组元素相同即可，基数统计
	numCnt := map[int]int{}
	for _, n := range target {
		numCnt[n] += 1
	}
	for _, n := range arr {
		numCnt[n] -= 1
	}
	for _, cnt := range numCnt {
		if cnt != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
