package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math/bits"
	"testing"
)

func Test_givenNormal_when762_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, countPrimeSetBits(6, 10))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countPrimeSetBits(left int, right int) (ans int) {
	// 32 以内的 prime
	// 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31
	const primeCnt = 0b010100000100010100010100010101100
	for i := left; i <= right; i++ {
		if primeCnt>>bits.OnesCount(uint(i))&0x01 == 1 {
			ans += 1
		}
	}
	return
}

func countPrimeSetBits1(left int, right int) int {
	// left <= right
	// n = 1的数量的最大值
	dp := [33]int{}
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		if dp[i] == 1 {
			continue
		}
		for j := 2; j*i < 33; j++ {
			dp[i*j] = 1
		}
	}
	ans := 0
	for i := left; i <= right; i++ {
		bit1 := 0
		for n := i; n != 0; n &= n - 1 {
			bit1 += 1
		}
		if dp[bit1] == 0 {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
