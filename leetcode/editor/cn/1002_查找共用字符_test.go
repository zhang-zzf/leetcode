package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1002_thenSuccess(t *testing.T) {
	assert.Equal(t, []string{"e", "l", "l"}, commonChars([]string{"bella", "label", "roller"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案后优化
func commonChars(words []string) []string {
	const charNum = 26
	dp := [charNum]int{}
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	for _, w := range words {
		cc := [charNum]int{}
		for _, c := range w {
			cc[c-'a'] += 1
		}
		for i := 0; i < len(cc); i++ {
			if dp[i] > cc[i] {
				dp[i] = cc[i]
			}
		}
	}
	var ans []string
	for i := 0; i < charNum; i++ {
		if dp[i] == 0 {
			continue
		}
		duplicate := string(rune('a' + i))
		for k := 0; k < dp[i]; k++ {
			ans = append(ans, duplicate)
		}
	}
	return ans
}

func commonChars1(words []string) []string {
	n := len(words)
	dp := make([][26]int, n)
	for i, w := range words {
		for _, c := range w {
			dp[i][c-'a'] += 1
		}
	}
	var ans []string
	for i := 0; i < 26; i++ {
		minCnt := math.MaxInt32
		for j := 0; j < n; j++ {
			if dp[j][i] < minCnt {
				minCnt = dp[j][i]
			}
		}
		duplicate := string([]byte{byte('a' + i)})
		for k := 0; k < minCnt; k++ {
			ans = append(ans, duplicate)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
