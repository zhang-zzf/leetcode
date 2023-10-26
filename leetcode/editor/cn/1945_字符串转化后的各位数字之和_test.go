package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1945_thenSuccess(t *testing.T) {
	assert.Equal(t, 8, getLucky("zbax", 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getLucky(s string, k int) int {
	toNum := func(x int) int {
		ans := 0
		for ; x > 0; x /= 10 {
			ans += x % 10
		}
		return ans
	}
	ans := 0
	for _, c := range s {
		position := int(c - 'a' + 1)
		ans += toNum(position)
	}
	for ; k > 1; k-- {
		ans = toNum(ans)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
