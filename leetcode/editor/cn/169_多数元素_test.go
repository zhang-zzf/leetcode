package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when169_thenSuccess(t *testing.T) {
	ans := majorityElement([]int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7})
	assert.Equal(t, 7, ans, "shouldEqual")
}

func Test_givenFailedCase1_when169_thenSuccess(t *testing.T) {
	ans := majorityElement([]int{3, 2, 3})
	assert.Equal(t, 3, ans, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// Boyer-Moore 投票算法
	ans, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 {
			ans = v
			cnt += 1
		} else if ans == v {
			cnt += 1
		} else {
			cnt -= 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
