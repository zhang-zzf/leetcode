package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when997_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, findJudge(2, [][]int{{1, 2}}))
	assert.Equal(t, 3, findJudge(3, [][]int{{1, 3}, {2, 3}}))
	assert.Equal(t, -1, findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}

func Test_givenFailedCase1_when997_thenSuccess(t *testing.T) {
	assert.Equal(t, 3, findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findJudge(n int, trust [][]int) int {
	dp := make([]int, 2*n+1)
	for _, t := range trust {
		dp[t[1]] += 1
		dp[n+t[0]] = 1
	}
	judge := -1
	for i := 1; i <= n; i++ {
		if dp[i] == n-1 && dp[n+i] == 0 {
			judge = i
		}
	}
	return judge
}

// 算法错误，一个人可以信任多个人
// [[1,3],[1,4],[2,3],[2,4],[4,3]]
func findJudge1(n int, trust [][]int) int {
	ans := -1
	if len(trust) == 0 {
		return ans
	}
	mayBeJudge, trustCnt := trust[0][1], 0
	for _, t := range trust {
		if t[1] == mayBeJudge {
			trustCnt += 1
			continue
		}
		// judge trust other
		if t[0] == mayBeJudge {
			mayBeJudge = -1
		}
		break
	}
	if trustCnt == n-1 {
		ans = mayBeJudge
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
