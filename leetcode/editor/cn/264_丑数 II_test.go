package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when264_thenSuccess(t *testing.T) {
	assert.Equal(t, 12, nthUglyNumber(10))
}

//leetcode submit region begin(Prohibit modification and deletion)
type Item struct {
	nth    int
	factor int
	ugly   int
}

func nthUglyNumber(n int) int {
	// TODO PriorityQueue
	dp := make([]int, n+1)
	pq := []*Item{
		{1, 2, 2},
		{1, 3, 2},
		{1, 5, 2},
	}
	for i := 2; i <= n; i++ {
		min := nthUglyNumberMin(pq, dp)
		if min == dp[len(dp)-1] {
			i -= 1
			continue
		}
		dp = append(dp, min)
	}
	return dp[n]
}

func nthUglyNumberMin(pq []*Item, dp []int) int {
	ans := pq[0]
	for i := 0; i < len(pq); i++ {
		pq[i].ugly = dp[pq[i].nth] * pq[i].factor
		if pq[i].ugly < ans.ugly {
			ans = pq[i]
		}
	}
	// update the pq
	ans.nth += 1
	return ans.ugly
}

//leetcode submit region end(Prohibit modification and deletion)
