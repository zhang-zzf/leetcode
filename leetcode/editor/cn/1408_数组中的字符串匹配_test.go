package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1408_thenSuccess(t *testing.T) {
	ans := stringMatching([]string{"mass", "as", "hero", "superhero"})
	assert.Equal(t, []string{"as", "hero"}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func stringMatching(words []string) []string {
	var ans []string
	for _, sub := range words {
		for _, w := range words {
			if sub != w && strings.Contains(w, sub) {
				ans = append(ans, sub)
				break
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
