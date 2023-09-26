package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1417_thenSuccess(t *testing.T) {
	ans := reformat("a0b1c2")
	assert.Equal(t, "0a1b2c", ans)
}

func Test_givenFailedCase1_when1417_thenSuccess(t *testing.T) {
	ans := reformat("covid2019")
	assert.Equal(t, "c2o0v1i9d", ans)
}

func Test_givenFailedCase_when1417_thenSuccess(t *testing.T) {
	ans := reformat("covid219")
	assert.Equal(t, "", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)

func reformat(s string) string {
	var ans, a1, a2 []byte
	for _, c := range []byte(s) {
		if c >= 'a' && c <= 'z' {
			a2 = append(a2, c)
		} else {
			a1 = append(a1, c)
		}
	}
	// failed case 1 covid2019
	// 数字字母的数量可以相差1个，第一位必须是较多数量的类型必须
	if len(a1) < len(a2) {
		a1, a2 = a2, a1
	}
	if len(a1)-len(a2) < 2 {
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				ans = append(ans, a1[i/2])
			} else {
				ans = append(ans, a2[i/2])
			}
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
