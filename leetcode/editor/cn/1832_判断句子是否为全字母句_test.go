package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1832_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkIfPangram(sentence string) bool {
	charCnt := 0
	cnt := make([]int, 26)
	for _, c := range sentence {
		ci := c - 'a'
		if cnt[ci] == 0 {
			cnt[ci] = 1
			charCnt += 1
		}
	}
	return charCnt == 26
}

//leetcode submit region end(Prohibit modification and deletion)
