package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when657_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func judgeCircle(moves string) bool {
	bytes := []byte(moves)
	x, y := 0, 0
	for _, c := range bytes {
		switch c {
		case 'L':
			x += 1
		case 'R':
			x -= 1
		case 'U':
			y += 1
		case 'D':
			y -= 1
		}
	}
	return x == 0 && y == 0
}

//leetcode submit region end(Prohibit modification and deletion)
