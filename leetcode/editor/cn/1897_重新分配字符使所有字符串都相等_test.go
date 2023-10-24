package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1897_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func makeEqual(words []string) bool {
	// TODO 数组比 map 性能高
	charCnt := make([]int, 26)
	for _, w := range words {
		for _, c := range []byte(w) {
			charCnt[c-'a'] += 1
		}
	}
	n := len(words)
	for _, cnt := range charCnt {
		if cnt%n != 0 {
			return false
		}
	}
	return true
}

func makeEqual1(words []string) bool {
	charCnt := map[byte]int{}
	for _, w := range words {
		for _, c := range []byte(w) {
			charCnt[c] += 1
		}
	}
	n := len(words)
	for _, cnt := range charCnt {
		if cnt%n != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
