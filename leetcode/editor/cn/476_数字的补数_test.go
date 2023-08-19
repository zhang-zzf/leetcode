package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when476_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, findComplement(5))
	assert.Equal(t, 0, findComplement(1))
	assert.Equal(t, 1, findComplement(math.MaxInt32-1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findComplement(num int) int {
	// 规律
	// 1. 5 101 -> 010 补数一定比 num 小
	// 2. num ^ 补数 == 111 = 2**3 - 1
	// 结论 补数 = (2**3 -1) ^ num
	n := int64(1)
	for n <= int64(num) {
		n <<= 1
	}
	return num ^ int(n-1)
}

func findComplement1(num int) int {
	// Time Limit Exceeded
	// 规律
	// 1. 5 101 -> 010 补数一定比 num 小
	// 2. num ^ 补数 == 111 = 2**3 - 1
	n := 1
	for n <= num {
		n *= 2
	}
	n = n - 1
	ans := 0
	for i := 1; i < num; i++ {
		if i^num == n {
			ans = i
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
