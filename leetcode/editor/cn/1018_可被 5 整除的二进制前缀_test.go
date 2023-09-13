package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1018_thenSuccess(t *testing.T) {
	ans := prefixesDivBy5([]int{1, 0, 1, 0, 1})
	assert.Equal(t, []bool{false, false, true, true, false}, ans)
}

func Test_givenFailedCase1_when1018_thenSuccess(t *testing.T) {
	ans := prefixesDivBy5([]int{0, 1, 1, 1, 1, 1})
	assert.Equal(t, []bool{true, false, false, false, true, false}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	mod5 := 0
	for i, n := range nums {
		// mod5 = (mod5*2 + nums[i]) % 5
		mod5 = (mod5<<1 | n) % 5
		ans[i] = mod5 == 0
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
