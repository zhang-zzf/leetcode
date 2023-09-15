package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1175_thenSuccess(t *testing.T) {
	assert.Equal(t, 682289015, numPrimeArrangements(100))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numPrimeArrangements(n int) int {
	const mod = 1e9 + 7
	factorial := func(f int) int {
		ans := 1
		for i := 1; i <= f; i++ {
			ans *= i
			if ans > mod {
				ans = ans % mod
			}
		}
		return ans
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		if dp[i] == 1 {
			continue
		}
		for j := 2; j < n+1; j++ {
			if i*j > n {
				break
			}
			dp[i*j] = 1
		}
	}
	primeCnt := 0
	for i := 2; i < n+1; i++ {
		if dp[i] == 0 {
			primeCnt += 1
		}
	}
	return factorial(primeCnt) * factorial(n-primeCnt) % mod
}

//leetcode submit region end(Prohibit modification and deletion)
