package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func Test_givenNormal_when1051_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func heightChecker(heights []int) (ans int) {
	// 基数排序
	max := math.MinInt
	for _, n := range heights {
		if n > max {
			max = n
		}
	}
	bucket := make([]int, max+1)
	for _, n := range heights {
		bucket[n] += 1
	}
	idx := 0
	for h, cnt := range bucket {
		for ; cnt > 0; cnt-- {
			if heights[idx] != h {
				ans += 1
			}
			idx += 1
		}
	}
	return ans
}
func heightChecker1(heights []int) (ans int) {
	// TODO 创建数组副本
	expected := append([]int{}, heights...)
	sort.Ints(expected)
	for i, v := range heights {
		if expected[i] != v {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
