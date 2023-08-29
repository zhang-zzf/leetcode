package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when746_thenSuccess(t *testing.T) {
	ans := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	assert.Equal(t, 6, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)

func minCostClimbingStairs(cost []int) int {
	// 动态规划
	n := len(cost)
	dp2, dp1 := 0, 0
	for i := 2; i < n+1; i++ {
		n1, n2 := dp1+cost[i-1], dp2+cost[i-2]
		dp2 = dp1
		dp1 = n1
		if n2 < dp1 {
			dp1 = n2
		}
	}
	return dp1
}

func minCostClimbingStairs1(cost []int) int {
	// 动态规划
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i < n+1; i++ {
		n1, n2 := dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]
		dp[i] = n1
		if n2 < dp[i] {
			dp[i] = n2
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
