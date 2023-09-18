package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func Test_givenNormal_when1200_thenSuccess(t *testing.T) {
	ans := minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27})
	assert.Equal(t, [][]int{{-14, -10}, {19, 23}, {23, 27}}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var ans [][]int
	subtract := math.MaxInt32
	for i := 1; i < len(arr); i++ {
		s := arr[i] - arr[i-1]
		if s < subtract {
			subtract = s
			ans = [][]int{{arr[i-1], arr[i]}}
		} else if s == subtract {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
