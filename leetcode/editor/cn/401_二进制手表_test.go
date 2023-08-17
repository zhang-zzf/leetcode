package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when401_thenSuccess(t *testing.T) {
	expected := []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}
	assert.Equal(t, expected, readBinaryWatch(1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 8 {
		return []string{}
	}
	// 时钟 [0,3] 个灯; 分钟 [0,5] 个灯
	// dp 记录 1 的个数
	dp := [60]int{}
	for i := 1; i < 60; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	// 计算出亮 【0-5】个灯的数字时
	cntPos := [6][]int{}
	for i := 0; i < 60; i++ {
		cntPos[dp[i]] = append(cntPos[dp[i]], i)
	}
	var ans []string
	for i := 0; i <= 3; i++ {
		j := turnedOn - i
		if j < 0 || j > 5 {
			continue
		}
		for _, h := range cntPos[i] {
			if h > 11 {
				continue
			}
			hour := strconv.Itoa(h)
			for _, m := range cntPos[j] {
				minute := strconv.Itoa(m)
				if m < 10 {
					minute = "0" + minute
				}
				ans = append(ans, hour+":"+minute)
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
