package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1742_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countBalls(lowLimit int, highLimit int) int {
	digitSum := func(x int) int {
		ans := 0
		for ; x > 0; x /= 10 {
			ans += x % 10
		}
		return ans
	}
	ans := 0
	cnt := make([]int, 46)
	for i := lowLimit; i <= highLimit; i++ {
		s := digitSum(i)
		cnt[s] += 1
		if cnt[s] > ans {
			ans = cnt[s]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
