package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1619_thenSuccess(t *testing.T) {
	ans := trimMean([]int{4, 8, 4, 10, 0, 7, 1, 3, 7, 8, 8, 3, 4, 1, 6, 2, 1, 1, 8, 0, 9, 8, 0, 3, 9, 10, 3, 10, 1, 10, 7, 3, 2, 1, 4, 9, 10, 7, 6, 4, 0, 8, 5, 1, 2, 1, 6, 2, 5, 0, 7, 10, 9, 10, 3, 7, 10, 5, 8, 5, 7, 6, 7, 6, 10, 9, 5, 10, 5, 5, 7, 2, 10, 7, 7, 8, 2, 0, 1, 1})
	assert.InDelta(t, 5.29167, ans, 0.00001)
}

//leetcode submit region begin(Prohibit modification and deletion)
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0
	for _, x := range arr[n/20 : n-n/20] {
		sum += x
	}
	return float64(sum*10) / float64(n*9)

}

//leetcode submit region end(Prohibit modification and deletion)
