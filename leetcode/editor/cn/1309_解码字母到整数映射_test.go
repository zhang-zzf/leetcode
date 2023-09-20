package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_givenNormal_when1309_thenSuccess(t *testing.T) {
	assert.Equal(t, "jkab", freqAlphabets("10#11#12"))
	assert.Equal(t, "acz", freqAlphabets("1326#"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func freqAlphabets(s string) string {
	// a b c d e f g h i j
	// 1 2 3 4 5 6 7 8 9 10
	ans := strings.Builder{}
	n := len(s)
	for i := 0; i < n; i++ {
		if i+2 < n && s[i+2] == '#' {
			idx, _ := strconv.Atoi(s[i : i+2])
			ans.WriteRune(rune('a' + idx - 1))
			i += 2
		} else {
			idx, _ := strconv.Atoi(s[i : i+1])
			ans.WriteRune(rune('a' + idx - 1))
		}
	}
	return ans.String()
}

//leetcode submit region end(Prohibit modification and deletion)
