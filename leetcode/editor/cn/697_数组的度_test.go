package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when697_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, findShortestSubArray([]int{1, 2, 2, 3, 1}))
	assert.Equal(t, 6, findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findShortestSubArray(nums []int) int {
	d, m := 1, 1
	mmap := map[int]*[2]int{}
	for i, n := range nums {
		if t, ok := mmap[n]; ok {
			t[0] += 1
			l := i - t[1] + 1
			if t[0] > d {
				d = t[0]
				m = l
			}
			if t[0] == d && l < m {
				m = l
			}
		} else {
			mmap[n] = &[2]int{1, i}
		}
	}
	return m
}

//leetcode submit region end(Prohibit modification and deletion)
