package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1790_thenSuccess(t *testing.T) {
	assert.Equal(t, true, areAlmostEqual("bank", "kanb"))
	assert.Equal(t, false, areAlmostEqual("attack", "defend"))
}

/**
边界
aa ac
*/
func Test_givenFailedCase1_when1790_thenSuccess(t *testing.T) {
	assert.Equal(t, false, areAlmostEqual("aa", "ac"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func areAlmostEqual(s1 string, s2 string) bool {
	cnt := 0
	for l, r := 0, len(s1)-1; l <= r; {
		//for l, r := 0, len(s1)-1; l < r; {
		if s1[l] == s2[l] {
			l += 1
			continue
		}
		if s1[r] == s2[r] {
			r -= 1
			continue
		}
		if cnt >= 1 || s1[l] != s2[r] || s1[r] != s2[l] {
			return false
		}
		l, r = l+1, r-1
		cnt += 1
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
