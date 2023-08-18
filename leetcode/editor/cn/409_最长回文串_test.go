package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when409_thenSuccess(t *testing.T) {
	assert.Equal(t, 7, longestPalindrome("abccccdd"))
	assert.Equal(t, 1, longestPalindrome("a"))
	assert.Equal(t, 7, longestPalindrome("aaaaaccc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	cntMap := map[rune]int{}
	for _, char := range s {
		cntMap[char] += 1
	}
	ans := 0
	for _, cnt := range cntMap {
		if ans%2 != 0 && cnt%2 != 0 {
			ans += cnt - 1
		} else {
			ans += cnt
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
