package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1389_thenSuccess(t *testing.T) {
	ans := createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})
	assert.Equal(t, []int{0, 4, 1, 3, 2}, ans)
}

func Test_givenFailedCase1_when1389_thenSuccess(t *testing.T) {
	ans := createTargetArray([]int{4, 2, 4, 3, 2}, []int{0, 0, 1, 3, 1})
	assert.Equal(t, []int{2, 2, 4, 4, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func createTargetArray(nums []int, index []int) []int {
	idxToNum := map[int]int{}
	var idx func(map[int]int, int, int)
	idx = func(m map[int]int, i, n int) {
		if en, ok := m[i]; ok {
			idx(m, i+1, en)
		}
		m[i] = n
	}
	for i, n := range nums {
		idx(idxToNum, index[i], n)
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = idxToNum[i]
	}
	return ans
}

// 算法是错误的
func createTargetArrayError(nums []int, index []int) []int {
	ans := make([]int, len(nums))
	for i, n := range index {
		if n == i {
			continue
		}
		for j := n; j < i; j++ {
			index[j] += 1
		}
	}
	for i, n := range nums {
		ans[index[i]] = n
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
