package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1736_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

/**
没有考虑后置
*/
func Test_givenFailedCase1_when1736_thenSuccess(t *testing.T) {
	assert.Equal(t, "14:03", maximumTime("?4:03"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumTime(time string) string {
	var ans []byte
	for i, c := range []byte(time) {
		if c != '?' {
			ans = append(ans, c)
			continue
		}
		switch {
		case i == 0 && time[1] >= '4' && time[1] <= '9':
			ans = append(ans, '1')
		case i == 0:
			ans = append(ans, '2')
		case i == 1 && ans[0] == '2':
			ans = append(ans, '3')
		case i == 1:
			ans = append(ans, '9')
		case i == 3:
			ans = append(ans, '5')
		case i == 4:
			ans = append(ans, '9')
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
