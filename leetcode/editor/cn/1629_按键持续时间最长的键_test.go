package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1629_thenSuccess(t *testing.T) {
	ans := slowestKey([]int{9, 29, 49, 50}, "cbcd")
	assert.Equal(t, byte('c'), ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func slowestKey(releaseTimes []int, keysPressed string) (ans byte) {
	keyDown, max := 0, 0
	for i, release := range releaseTimes {
		time := release - keyDown
		if time > max {
			max = time
			ans = keysPressed[i]
		} else if time == max && keysPressed[i] > ans {
			ans = keysPressed[i]
		}
		keyDown = release
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
