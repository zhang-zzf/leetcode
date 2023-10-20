package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1812_thenSuccess(t *testing.T) {
	assert.Equal(t, false, squareIsWhite("a1"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func squareIsWhite(coordinates string) bool {
	return (coordinates[0]-'a'+coordinates[1]-'1')%2 == 1
}

//leetcode submit region end(Prohibit modification and deletion)
