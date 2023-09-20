package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1317_thenSuccess(t *testing.T) {
	ans := getNoZeroIntegers(1010)
	assert.Equal(t, []int{11, 999}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func getNoZeroIntegers(n int) []int {
	var ans []int
	containsZero := func(a int) bool {
		ans := false
		for a > 0 {
			if a%10 == 0 {
				ans = true
				break
			}
			a /= 10
		}
		return ans
	}
	for i := 1; i < n; i++ {
		if containsZero(i) {
			continue
		}
		o := n - i
		if containsZero(o) {
			continue
		}
		ans = append(ans, i, o)
		break
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
