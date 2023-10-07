package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1528_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func restoreString(s string, indices []int) string {
	bytes := make([]byte, len(s))
	for i, c := range []byte(s) {
		bytes[indices[i]] = c
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
