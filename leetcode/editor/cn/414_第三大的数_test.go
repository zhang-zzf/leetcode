package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when414_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, thirdMax([]int{3, 2, 1}))
	assert.Equal(t, 3, thirdMax([]int{3}))
	assert.Equal(t, 3, thirdMax([]int{3, 2}))
	assert.Equal(t, math.MinInt32, thirdMax([]int{math.MinInt32}))
}

/**
如果不存在，则返回数组中最大的数。
*/
func Test_givenFailedCase1_when414_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, thirdMax([]int{3, 2}))
}

/**
[2,2,3,1]
*/
func Test_givenFailedCase2_when414_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, thirdMax([]int{2, 2, 3, 1}))
}

/**
[1,1,2]
*/
func Test_givenFailedCase3_when414_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, thirdMax([]int{1, 1, 2}))
}

/**
[1,2,2,5,3,5]
*/
func Test_givenFailedCaser_when414_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, thirdMax([]int{1, 2, 2, 5, 3, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n == m1 || n == m2 || n == m3 {
			continue
		}
		if n > m1 {
			m1, m2, m3 = n, m1, m2
		} else if n > m2 {
			m2, m3 = n, m2
		} else if n != m2 && n > m3 {
			m3 = n
		}
	}
	if m3 == math.MinInt64 {
		return m1
	} else {
		return m3
	}
}

//leetcode submit region end(Prohibit modification and deletion)
