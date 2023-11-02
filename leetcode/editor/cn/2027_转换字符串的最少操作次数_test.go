package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2027_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, minimumMoves("OOX"))
	assert.Equal(t, 1, minimumMoves("XOX"))
	assert.Equal(t, 1, minimumMoves("OXOX"))
	assert.Equal(t, 2, minimumMoves("XXOX"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(s string) int {
	ans := 0
	xCnt := 0
	for _, c := range s {
		if c == 'X' || xCnt > 0 {
			xCnt += 1
		}
		if xCnt == 3 {
			ans += 1
			xCnt = 0
		}
	}
	if xCnt > 0 {
		ans += 1
	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
