package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1108_thenSuccess(t *testing.T) {
	assert.Equal(t, "255[.]100[.]50[.]0", defangIPaddr("255.100.50.0"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

//leetcode submit region end(Prohibit modification and deletion)
