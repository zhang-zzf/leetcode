package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when459_thenSuccess(t *testing.T) {
	assert.Equal(t, true, repeatedSubstringPattern("ababab"))
	assert.Equal(t, true, repeatedSubstringPattern("aa"))
	assert.Equal(t, true, repeatedSubstringPattern("aaa"))
	assert.Equal(t, false, repeatedSubstringPattern("ab"))
}

func Test_givenFailedCase1_when459_thenSuccess(t *testing.T) {
	assert.Equal(t, true, repeatedSubstringPattern("abaababaab"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 在 repeatedSubstringPattern1 基础上优化
func repeatedSubstringPattern(s string) bool {
	ans := false
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		ans = true
		for j := i; j < n; j++ {
			if s[j] != s[j-i] {
				ans = false
				break
			}
		}
		if ans {
			break
		}
	}
	return ans
}
func repeatedSubstringPattern1(s string) bool {
	// 暴力破解
	ans := false
	n := len(s)
	for subLen := 1; subLen <= n/2; subLen++ {
		if n%subLen != 0 {
			continue
		}
		ans = true
		repeat := n / subLen
		for j := 0; j < subLen; j++ {
			for k := 0; k < repeat; k++ {
				if s[j] != s[j+k*subLen] {
					ans = false
					break
				}
			}
		}
		if ans {
			break
		}
	}
	return ans
}

// error "ababab"
// "aaa"
func repeatedSubstringPatternError(s string) bool {
	// r 向上取整
	n := len(s)
	if n%2 != 0 {
		return false
	}
	l, r := 0, n/2
	ans := true
	for r < n {
		if s[l] != s[r] {
			ans = false
		}
		l, r = l+1, r+1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
