package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1160_thenSuccess(t *testing.T) {
	ans := countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach")
	assert.Equal(t, 6, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案后
func countCharacters(words []string, chars string) (ans int) {
	charCnt := make([]int, 26)
	for _, c := range chars {
		charCnt[c-'a'] += 1
	}
	for _, w := range words {
		cc := append([]int{}, charCnt...)
		canSpell := true
		for _, c := range w {
			if cc[c-'a'] -= 1; cc[c-'a'] < 0 {
				canSpell = false
				break
			}
		}
		if canSpell {
			ans += len(w)
		}
	}
	return ans
}
func countCharacters1(words []string, chars string) int {
	ans := 0
	canSpell := func(w, cs string) bool {
		charCnt := map[byte]int{}
		for _, c := range []byte(cs) {
			charCnt[c] += 1
		}
		for _, c := range []byte(w) {
			if charCnt[c] -= 1; charCnt[c] < 0 {
				return false
			}
		}
		return true
	}
	for _, w := range words {
		if canSpell(w, chars) {
			ans += len(w)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
