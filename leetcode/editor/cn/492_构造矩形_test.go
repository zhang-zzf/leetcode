package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when492_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func constructRectangle(area int) []int {
	l, w := 0, int(math.Sqrt(float64(area)))
	for w > 0 {
		if area%w == 0 {
			l = area / w
			break
		}
		w -= 1
	}
	return []int{l, w}
}

//leetcode submit region end(Prohibit modification and deletion)
