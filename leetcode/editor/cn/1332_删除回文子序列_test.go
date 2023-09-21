package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1332_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, removePalindromeSub("ababa"))
	assert.Equal(t, 1, removePalindromeSub("abaaba"))
	assert.Equal(t, 2, removePalindromeSub("ababba"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func removePalindromeSub(s string) int {
	// 分析后答案只可能是 1 / 2
	l, r := 0, len(s)-1
	for ; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			break
		}
	}
	if l >= r { // 本身是回文串
		return 1
	}
	return 2
}

//leetcode submit region end(Prohibit modification and deletion)
