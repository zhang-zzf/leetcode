package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2432_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func hardestWorker(n int, logs [][]int) int {
	worker, startTime, longestTask := 0, 0, 0
	for _, log := range logs {
		task := log[1] - startTime
		if task > longestTask {
			worker = log[0]
			longestTask = task
		} else if task == longestTask && log[0] < worker {
			worker = log[0]
		}
		startTime = log[1]
	}
	return worker
}

//leetcode submit region end(Prohibit modification and deletion)
