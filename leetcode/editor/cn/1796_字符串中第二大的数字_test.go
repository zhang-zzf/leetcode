package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1796_thenSuccess(t *testing.T) {
	assert.Equal(t, -1, secondHighest("abc1111"))
	assert.Equal(t, 1, secondHighest("abc1112"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func secondHighest(s string) int {
	m1, m2 := -1, -1
	for _, c := range s {
		if c <= '9' {
			n := int(c - '0')
			if n > m1 {
				m1, m2 = n, m1
			} else if n != m1 && n > m2 {
				m2 = n
			}
		}
	}
	return m2
}

//leetcode submit region end(Prohibit modification and deletion)
