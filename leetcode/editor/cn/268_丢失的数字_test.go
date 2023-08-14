package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when268_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{3, 0, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	arrSum := 0
	for _, val := range nums {
		arrSum += val
	}
	return sum - arrSum
}

//leetcode submit region end(Prohibit modification and deletion)
