package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when205_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	ans := true
	m1, m2 := map[byte]byte{}, map[byte]byte{}
	for idx := 0; idx < len(s); idx++ {
		sc, tc := s[idx], t[idx]
		if mt, exists := m1[sc]; !exists {
			if ms, exists := m2[tc]; !exists || ms == sc {
				m1[sc], m2[tc] = tc, sc
			} else {
				ans = false
				break
			}
		} else if mt != tc {
			ans = false
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
