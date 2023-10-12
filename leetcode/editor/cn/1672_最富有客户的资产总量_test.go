package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1672_thenSuccess(t *testing.T) {
	ans := maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})
	assert.Equal(t, 17, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumWealth(accounts [][]int) (ans int) {
	sum := func(arr ...int) (ans int) {
		for _, x := range arr {
			ans += x
		}
		return ans
	}
	for _, assets := range accounts {
		total := sum(assets...)
		if total > ans {
			ans = total
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
