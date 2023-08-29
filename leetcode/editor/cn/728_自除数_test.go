package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when728_thenSuccess(t *testing.T) {
	ans := selfDividingNumbers(1, 22)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func selfDividingNumbers(left int, right int) []int {
	var isSelfDividing func(int) bool
	isSelfDividing = func(nn int) bool {
		ans := true
		for n := nn; n > 0; n /= 10 {
			mod := n % 10
			if mod == 0 || nn%mod != 0 {
				ans = false
				break
			}
		}
		return ans
	}
	var ans []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
