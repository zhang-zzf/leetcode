package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1925_thenSuccess(t *testing.T) {
	assert.Equal(t, 10, countTriples(18))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countTriples(n int) (ans int) {
	// TODO 暴力破解
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			sqrSum := i*i + j*j
			sqr := int(math.Sqrt(float64(sqrSum)))
			if sqr <= n && sqr*sqr == sqrSum {
				ans += 1
			}
		}
	}
	return ans * 2
}

//leetcode submit region end(Prohibit modification and deletion)
