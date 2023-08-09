package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when228_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func summaryRanges(nums []int) []string {
	var ans []string
	for left, right := 0, 0; right < len(nums); right += 1 {
		if right == len(nums)-1 || nums[right+1]-nums[right] > 1 {
			numRange := strconv.Itoa(nums[left])
			if left != right {
				numRange += "->" + strconv.Itoa(nums[right])
			}
			ans = append(ans, numRange)
			left = right + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
