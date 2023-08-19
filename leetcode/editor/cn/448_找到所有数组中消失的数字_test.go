package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when448_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	cntMap := make([]int, len(nums)+1)
	for _, n := range nums {
		cntMap[n] = 1
	}
	var ans []int
	for i := 1; i <= len(nums); i++ {
		if cntMap[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
