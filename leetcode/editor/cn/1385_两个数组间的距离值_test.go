package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1385_thenSuccess(t *testing.T) {
	ans := findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)
	assert.Equal(t, 2, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) (ans int) {
	sort.Ints(arr2)
	n := len(arr2)
	for _, num := range arr1 {
		l, r := num-d, num+d
		idx := sort.Search(n, func(i int) bool { return arr2[i] >= num })
		if (idx == n || arr2[idx] > r) && (idx == 0 || arr2[idx-1] < l) {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
