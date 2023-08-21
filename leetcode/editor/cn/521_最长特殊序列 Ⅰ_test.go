package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when521_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, findLUSlength("aba", "cdc"))
	assert.Equal(t, 3, findLUSlength("ab", "cdc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(a string, b string) int {
	m, n := len(a), len(b)
	if m > n {
		m, n = n, m
	}
	if m != n {
		return n
	} else if a != b {
		return n
	} else {
		return -1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
