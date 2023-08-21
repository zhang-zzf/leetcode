package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when504_thenSuccess(t *testing.T) {
	assert.Equal(t, "202", convertToBase7(100))
	assert.Equal(t, "-10", convertToBase7(-7))
}

/**
边界条件 0
*/
func Test_givenFailedCase1_when504_thenSuccess(t *testing.T) {
	assert.Equal(t, "0", convertToBase7(0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var bytes []byte
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	for num != 0 {
		bytes = append(bytes, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		bytes = append(bytes, '-')
	}
	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
