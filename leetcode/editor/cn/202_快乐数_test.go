package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when202_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isHappy(19))
	assert.Equal(t, false, isHappy(2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	// 不快乐数唯一循环序列 4 16 37 58 89 145 42 20
	// 若 n 的不快乐数是上述序列中的任意一个值，n 便是不快乐数
	ans := false
	for true {
		n = nextN(n)
		if n == 1 {
			ans = true
			break
		} else if n == 4 {
			break
		}
	}
	return ans

}
func nextN(n int) int {
	ans := 0
	for n != 0 {
		mod := n % 10
		ans += mod * mod
		n /= 10
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
