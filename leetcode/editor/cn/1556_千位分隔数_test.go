package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1556_thenSuccess(t *testing.T) {
	assert.Equal(t, "987", thousandSeparator(987))
	assert.Equal(t, "1.234", thousandSeparator(1234))
	assert.Equal(t, "123.456", thousandSeparator(123456))
}

/**
边界条件0
*/
func Test_givenFailedCase1_when1556_thenSuccess(t *testing.T) {
	assert.Equal(t, "0", thousandSeparator(0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	var ans []byte
	cnt := 0
	for n > 0 {
		if cnt == 3 {
			ans = append(ans, '.')
			cnt = 0
		}
		bit := byte(n % 10)
		ans = append(ans, '0'+bit)
		cnt += 1
		n /= 10
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
