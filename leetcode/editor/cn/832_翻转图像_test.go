package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when832_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		for l, r := 0, len(row)-1; l <= r; l, r = l+1, r-1 {
			row[l], row[r] = row[r]^1, row[l]^1
		}
	}
	return image
}

//leetcode submit region end(Prohibit modification and deletion)
