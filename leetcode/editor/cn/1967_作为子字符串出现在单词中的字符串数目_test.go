package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1967_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numOfStrings(patterns []string, word string) (ans int) {
	for _, w := range patterns {
		if strings.Contains(word, w) {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
