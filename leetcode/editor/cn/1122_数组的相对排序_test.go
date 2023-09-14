package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1122_thenSuccess(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	ans := relativeSortArray(arr1, arr2)
	assert.Equal(t, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var ans []int
	nCnt := map[int]int{}
	for _, n := range arr1 {
		nCnt[n] += 1
	}
	for _, n := range arr2 {
		for ; nCnt[n] > 0; nCnt[n]-- {
			ans = append(ans, n)
		}
	}
	sortCnt := 0
	for k, c := range nCnt {
		for ; c > 0; c-- {
			ans = append(ans, k)
			sortCnt += 1
		}
	}
	if sortCnt > 0 {
		sort.Ints(ans[len(arr1)-sortCnt:])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
