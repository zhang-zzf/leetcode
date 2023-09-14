package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1078_thenSuccess(t *testing.T) {
	ans := findOcurrences("we will we will rock you", "we", "will")
	assert.Equal(t, []string{"we", "rock"}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findOcurrences(text string, first string, second string) []string {
	var ans []string
	words := strings.Split(text, " ")
	for i, w := range words {
		if w == first && i+2 < len(words) && words[i+1] == second {
			ans = append(ans, words[i+2])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
