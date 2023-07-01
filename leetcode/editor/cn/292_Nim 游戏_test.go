package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when292_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canWinNim(1), "shouldEqual")
	assert.Equal(t, false, canWinNim(4), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func canWinNim(n int) bool {
	return (n % 4) != 0
}

//leetcode submit region end(Prohibit modification and deletion)
