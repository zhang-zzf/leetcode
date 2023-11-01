package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2000_thenSuccess(t *testing.T) {
	assert.Equal(t, "dcbaefd", reversePrefix("abcdefd", 'd'))
	assert.Equal(t, "abcd", reversePrefix("abcd", 'z'))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reversePrefix(word string, ch byte) string {
	ans := []byte(word)
	idx := 0
	for ; idx < len(ans); idx++ {
		if ans[idx] == ch {
			break
		}
	}
	if idx == len(ans) {
		return word
	}
	for l, r := 0, idx; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
