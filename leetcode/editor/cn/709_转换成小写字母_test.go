package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when709_thenSuccess(t *testing.T) {
	assert.Equal(t, "hello", toLowerCase("Hello"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func toLowerCase(s string) string {
	bytes := []byte(s)
	for i, c := range bytes {
		if c >= 'A' && c <= 'Z' {
			bytes[i] = c + 'a' - 'A'
		}
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
