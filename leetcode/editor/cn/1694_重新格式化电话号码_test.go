package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1694_thenSuccess(t *testing.T) {
	ans := reformatNumber("123 4-567")
	assert.Equal(t, "123-45-67", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reformatNumber(number string) string {
	var ans []byte
	cnt := 0
	for _, c := range []byte(number) {
		if c == ' ' || c == '-' {
			continue
		} else {
			if cnt == 3 {
				ans = append(ans, '-')
				cnt = 0
			}
			ans = append(ans, c)
			cnt += 1
		}
	}
	if cnt == 1 {
		n := len(ans)
		ans[n-3], ans[n-2] = ans[n-2], ans[n-3]
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
