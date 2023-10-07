package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1539_thenSuccess(t *testing.T) {
	ans := findKthPositive([]int{2, 3, 4, 7, 11}, 5)
	assert.Equal(t, 9, ans)
}

/**
边界条件
*/
func Test_givenFailedCase1_when1539_thenSuccess(t *testing.T) {
	ans := findKthPositive([]int{1, 2, 3, 4}, 2)
	assert.Equal(t, 6, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthPositive(arr []int, k int) int {
	// divide
	total := len(arr) + k
	numCnt := make([]int, total)
	cnt, idx := 0, 0
	for i := 1; i < total; i++ {
		if idx < len(arr) && i == arr[idx] {
			idx += 1
		} else {
			cnt += 1
		}
		numCnt[i] = cnt
	}
	return sort.Search(total, func(i int) bool { return numCnt[i] >= k })
}

// 暴力破
func findKthPositive1(arr []int, k int) int {
	n, idx := 1, 0
	for {
		// TODO idx 边界条件
		if idx >= len(arr) || n < arr[idx] {
			k -= 1
		} else {
			idx += 1
		}
		if k == 0 {
			break
		}
		n += 1
	}
	return n
}

//leetcode submit region end(Prohibit modification and deletion)
