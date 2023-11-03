package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2108_thenSuccess(t *testing.T) {
	ans := firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"})
	assert.Equal(t, "ada", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func firstPalindrome(words []string) string {
	isPalindrome := func(w string) bool {
		for l, r := 0, len(w)-1; l < r; l, r = l+1, r-1 {
			if w[l] != w[r] {
				return false
			}
		}
		return true
	}
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}
	return ""

}

//leetcode submit region end(Prohibit modification and deletion)
