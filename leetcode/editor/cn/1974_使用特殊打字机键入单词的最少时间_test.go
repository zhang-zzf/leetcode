package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1974_thenSuccess(t *testing.T) {
	assert.Equal(t, 34, minTimeToType("zjpc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minTimeToType(word string) int {
	ans := 0
	prev := 'a'
	for _, w := range word {
		distance := w - prev
		if distance < 0 {
			distance = -distance
		}
		if distance > 13 {
			distance = 26 - distance
		}
		ans += int(distance) + 1
		prev = w
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
