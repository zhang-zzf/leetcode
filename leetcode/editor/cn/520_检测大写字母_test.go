package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when520_thenSuccess(t *testing.T) {
	assert.Equal(t, true, detectCapitalUse("U"))
	assert.Equal(t, true, detectCapitalUse("USA"))
	assert.Equal(t, true, detectCapitalUse("Leetcode"))
	assert.Equal(t, true, detectCapitalUse("leet"))
	assert.Equal(t, true, detectCapitalUse("u"))
	assert.Equal(t, false, detectCapitalUse("lL"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	ans := true
	u := upper(word[0]) & upper(word[1])
	for i := 1; i < len(word); i++ {
		if u^upper(word[i]) != 0 {
			ans = false
			break
		}
	}
	return ans
}

func upper(w uint8) int {
	if w >= 'A' && w <= 'Z' {
		return 1
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
