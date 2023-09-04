package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when860_thenSuccess(t *testing.T) {
	assert.Equal(t, true, lemonadeChange([]int{5, 5, 5, 10, 20}))
	assert.Equal(t, false, lemonadeChange([]int{5, 5, 10, 10, 20}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lemonadeChange(bills []int) bool {
	deposits := [3]int{}
	ans := true
	for _, n := range bills {
		switch n {
		case 5:
			deposits[0] += 1
		case 10:
			deposits[1] += 1
			deposits[0] -= 1
		case 20:
			deposits[2] += 1
			if deposits[1] > 0 {
				deposits[1] -= 1
				deposits[0] -= 1
			} else {
				deposits[0] -= 3
			}
		}
		ans = deposits[0] >= 0 && deposits[1] >= 0
		if ans == false {
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
