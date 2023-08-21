package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when485_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
	cnt, maxCnt := 0, 0
	for _, n := range nums {
		if n == 1 {
			cnt += 1
			if cnt > maxCnt {
				maxCnt = cnt
			}
		} else {
			cnt = 0
		}
	}
	return maxCnt
}

//leetcode submit region end(Prohibit modification and deletion)
