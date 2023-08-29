package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when748_thenSuccess(t *testing.T) {
	ans := shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"})
	assert.Equal(t, "pest", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestCompletingWord(licensePlate string, words []string) string {
	lm := charMap(licensePlate)
	ans := ""
	for _, w := range words {
		m := charMap(w)
		match := true
		for k, v := range lm {
			if v > m[k] {
				match = false
				break
			}
		}
		if match && (ans == "" || len(ans) > len(w)) {
			ans = w
		}
	}
	return ans
}

func charMap(s string) map[byte]int {
	cntMap := map[byte]int{}
	for _, c := range []byte(s) {
		if c >= 'A' && c <= 'Z' {
			c = c + 'a' - 'A'
		}
		if c >= 'a' && c <= 'z' {
			cntMap[c] += 1
		}
	}
	return cntMap
}

//leetcode submit region end(Prohibit modification and deletion)
