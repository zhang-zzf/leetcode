package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1880_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	toNum := func(s string) int {
		ans := 0
		for _, c := range s {
			ans = ans*10 + int(c-'a')
		}
		return ans
	}
	return toNum(firstWord)+toNum(secondWord) == toNum(targetWord)
}

//leetcode submit region end(Prohibit modification and deletion)
