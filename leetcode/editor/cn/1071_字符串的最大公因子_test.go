package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1071_thenSuccess(t *testing.T) {
	assert.Equal(t, "ABC", gcdOfStrings("ABCABC", "ABC"))
	assert.Equal(t, "NLZGM", gcdOfStrings("NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM",
		"NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func gcdOfStrings(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	if m > n {
		m, n, str1, str2 = n, m, str2, str1
	}
	isGCD := func(gcd, str string) bool {
		lg, ls := len(gcd), len(str)
		if ls%lg != 0 {
			return false
		}
		for i := 0; i < ls; i++ {
			if (i+1)%lg != 0 {
				continue
			}
			if gcd != str[i-lg+1:i+1] {
				return false
			}
		}
		return true
	}
	// 分成几份
	for i := 1; i <= m; i++ {
		if m%i != 0 {
			continue
		}
		gcd := str1[0 : m/i]
		if isGCD(gcd, str1) && isGCD(gcd, str2) {
			return gcd
		}
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
