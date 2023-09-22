package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1356_thenSuccess(t *testing.T) {
	params := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	assert.Equal(t, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, sortByBits(params))
	params = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}, sortByBits(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortByBits(arr []int) []int {
	count1 := func(x int) int {
		ans := 0
		for x > 0 {
			x &= x - 1
			ans += 1
		}
		return ans
	}
	n := len(arr)
	ps := make([]pair1356, n)
	for i := 0; i < n; i++ {
		ps[i] = pair1356{count1(arr[i]), arr[i]}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].cnt < ps[j].cnt || (ps[i].cnt == ps[j].cnt && ps[i].num < ps[j].num)
	})
	for i := 0; i < n; i++ {
		arr[i] = ps[i].num
	}
	return arr
}

type pair1356 struct {
	cnt, num int
}

//leetcode submit region end(Prohibit modification and deletion)
