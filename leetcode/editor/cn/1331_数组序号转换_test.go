package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1331_thenSuccess(t *testing.T) {
	ans := arrayRankTransform([]int{40, 10, 20, 30})
	assert.Equal(t, []int{4, 1, 2, 3}, ans)
}

func Test_givenFailedCase1_when1331_thenSuccess(t *testing.T) {
	ans := arrayRankTransform([]int{1, 1, 1})
	assert.Equal(t, []int{1, 1, 1}, ans)
}

func Test_givenFailedCase2_when1331_thenSuccess(t *testing.T) {
	ans := arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12})
	assert.Equal(t, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// 在 1 的基础上性能优化
func arrayRankTransform(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	idx, numToIdx := 1, map[int]int{}
	for _, n := range arrCopy {
		if _, ok := numToIdx[n]; ok {
			continue
		}
		numToIdx[n] = idx
		idx += 1
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = numToIdx[arr[i]]
	}
	return arr
}

func arrayRankTransform1(arr []int) []int {
	var arrCopy []int
	arrCopy = append(arrCopy, arr...)
	sort.Ints(arrCopy)
	numToIdx := map[int]int{}
	sortNum := 1
	for _, n := range arrCopy {
		if _, exists := numToIdx[n]; !exists {
			numToIdx[n] = sortNum
			sortNum += 1
		}
	}
	var ans []int
	for _, n := range arr {
		ans = append(ans, numToIdx[n])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
