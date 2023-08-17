package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when389_thenSuccess(t *testing.T) {
	diff := findTheDifference("abcd", "adcbe")
	assert.Equal(t, byte('e'), diff)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	var cntMap [26]int
	for _, char := range s {
		cntMap[char-'a'] += 1
	}
	for _, char := range t {
		cntMap[char-'a'] -= 1
	}
	for idx, cnt := range cntMap {
		if cnt == -1 {
			return byte(idx) + 'a'
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
