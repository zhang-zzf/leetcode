package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when171_thenSuccess(t *testing.T) {
	number := titleToNumber("ZY")
	assert.Equal(t, 701, number)
	assert.Equal(t, 26, 26<<0)
	// 26<<1 = 26 * 2
	assert.Equal(t, 52, 26<<1)
	assert.NotEqual(t, int(math.Pow(26, 3)), 26<<2)
}

//leetcode submit region begin(Prohibit modification and deletion)
func titleToNumber(columnTitle string) int {
	ans := 0
	for _, c := range columnTitle {
		ans = ans*26 + int(c-'A'+1)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
