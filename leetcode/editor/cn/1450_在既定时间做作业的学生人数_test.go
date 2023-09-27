package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1450_thenSuccess(t *testing.T) {
	startTime := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	endTime := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	ans := busyStudent(startTime, endTime, 5)
	assert.Equal(t, 5, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ans := 0
	for i, start := range startTime {
		if start <= queryTime && endTime[i] >= queryTime {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
