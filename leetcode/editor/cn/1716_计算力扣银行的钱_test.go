package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1716_thenSuccess(t *testing.T) {
	assert.Equal(t, 37, totalMoney(10))
	assert.Equal(t, 96, totalMoney(20))
}

//leetcode submit region begin(Prohibit modification and deletion)
func totalMoney(n int) int {
	ans, m, w := 0, 1, 1
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			m = w
			w += 1
		}
		ans += m
		m += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
