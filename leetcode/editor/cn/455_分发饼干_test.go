package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when455_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	// 排序+双指针
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	ptrg, ptrs := 0, 0
	for ptrg < len(g) && ptrs < len(s) {
		if s[ptrs] < g[ptrg] {
			ptrs += 1
			continue
		}
		ans += 1
		ptrg, ptrs = ptrg+1, ptrs+1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
