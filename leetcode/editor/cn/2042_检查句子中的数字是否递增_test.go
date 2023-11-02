package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2042_thenSuccess(t *testing.T) {
	assert.Equal(t, true, areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	assert.Equal(t, false, areNumbersAscending("1 box has 3 blue 4 red 6 green and 3"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func areNumbersAscending(s string) bool {
	prevNum := 0
	num := 0
	for _, c := range []byte(s + " ") {
		if c == ' ' {
			if num == 0 {
				continue
			}
			if num <= prevNum {
				return false
			}
			prevNum = num
			num = 0
		} else if c <= '9' {
			num = num*10 + int(c-'0')
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
