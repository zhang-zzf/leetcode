package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1979_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, findGCD([]int{7, 5, 6, 8, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		} else if n < min {
			min = n
		}
	}
	var gcd func(int, int) int
	gcd = func(n1, n2 int) int {
		if n2 == 0 {
			return n1
		}
		return gcd(n2, n1%n2)
	}
	return gcd(min, max)
}

//leetcode submit region end(Prohibit modification and deletion)
