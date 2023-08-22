package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when594_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	assert.Equal(t, 2, findLHS([]int{1, 2, 3, 4}))
	assert.Equal(t, 0, findLHS([]int{1, 1, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findLHS(nums []int) int {
	cntMap := map[int]int{}
	for _, n := range nums {
		cntMap[n] += 1
	}
	ans := 0
	for k, v := range cntMap {
		if j, ok := cntMap[k+1]; ok {
			if j+v > ans {
				ans = j + v
			}
		}
	}
	return ans
}

func findLHS1(nums []int) int {
	cntMap := map[int]int{}
	for _, n := range nums {
		cntMap[n] += 1
	}
	ans := 0
	for _, n := range nums {
		m := 0
		if m1, ok := cntMap[n+1]; ok {
			m = cntMap[n] + m1
		}
		if m > ans {
			ans = m
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
