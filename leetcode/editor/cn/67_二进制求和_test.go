package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when67_thenSuccess(t *testing.T) {
	ans := addBinary("11", "111")
	assert.Equal(t, "1010", ans, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	ans := ""
	var promotion uint8
	for ia, ib := len(a)-1, len(b)-1; ia >= 0 || ib >= 0; ia, ib = ia-1, ib-1 {
		var na, nb uint8
		if ia >= 0 {
			na = a[ia] - '0'
		}
		if ib >= 0 {
			nb = b[ib] - '0'
		}
		sum := na + nb + promotion
		ans = strconv.FormatInt(int64(sum%2), 10) + ans
		promotion = sum / 2
	}
	if promotion > 0 {
		ans = "1" + ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
