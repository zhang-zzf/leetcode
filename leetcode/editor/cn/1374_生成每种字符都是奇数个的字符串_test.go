package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1374_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func generateTheString(n int) string {
	bytes := make([]byte, n)
	for i := 0; i < n-1; i++ {
		bytes[i] = 'a'
	}
	if (n-1)&0x01 == 0 {
		bytes[n-1] = 'a'
	} else {
		bytes[n-1] = 'b'
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
