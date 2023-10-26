package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenFailedCase1_when1961_thenSuccess(t *testing.T) {
	ans := isPrefixString("a", []string{"aa", "aaaa", "banana"})
	assert.Equal(t, false, ans)
}

func Test_givenFailedCase2_when1961_thenSuccess(t *testing.T) {
	ans := isPrefixString("aaa", []string{"aa", "aaaa", "banana"})
	assert.Equal(t, false, ans)
}

func Test_givenFailedCase3_when1961_thenSuccess(t *testing.T) {
	ans := isPrefixString("cccccc", []string{"c", "c"})
	assert.Equal(t, false, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPrefixString(s string, words []string) bool {
	aStr := ""
	ns, nw := len(s), len(words)
	// todo 边界条件
	// for i := 0; len(aStr) < n; i++ {
	for i := 0; i < nw && len(aStr) < ns; i++ {
		aStr += words[i]
	}
	return s == aStr
}

// 没有理解题意
func isPrefixStringError(s string, words []string) bool {
	// todo 边界
	if len(s) < len(words[0]) {
		return false
	}
	return strings.HasPrefix(strings.Join(words, ""), s)
}

//leetcode submit region end(Prohibit modification and deletion)
