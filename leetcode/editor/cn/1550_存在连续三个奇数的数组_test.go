package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1550_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeConsecutiveOdds(arr []int) bool {
	for l := 0; l < len(arr); l++ {
		if arr[l]&0x01 == 0 {
			continue
		}
		oddCnt := 1
		r := l + 1
		for ; r < len(arr); r++ {
			if arr[r]&0x01 == 0 || oddCnt >= 3 {
				break
			}
			oddCnt += 1
		}
		if oddCnt >= 3 {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
