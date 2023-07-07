package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_whenLCP17_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, calculate("AB"), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	x, y := 1, 0
	for _, c := range s {
		switch c {
		case 'A':
			x = x*2 + y
		case 'B':
			y = y*2 + x
		default:
		}
	}
	return x + y
}

//leetcode submit region end(Prohibit modification and deletion)
