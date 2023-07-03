package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when415_thenSuccess(t *testing.T) {
	assert.Equal(t, "10", addStrings("1", "9"), "shouldEqual")
	assert.Equal(t, "134", addStrings("11", "123"), "shouldEqual")
	assert.Equal(t, "0", addStrings("0", "0"), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	var ans []byte
	var promotion byte
	for i1, i2 := len(num1)-1, len(num2)-1; i1 >= 0 || i2 >= 0; i1, i2 = i1-1, i2-1 {
		var n1, n2 byte
		if i1 >= 0 {
			n1 = num1[i1] - '0'
		}
		if i2 >= 0 {
			n2 = num2[i2] - '0'
		}
		sum := n1 + n2 + promotion
		if sum < 10 {
			promotion = 0
		} else {
			promotion = 1
			sum -= 10
		}
		ans = append([]byte{sum + '0'}, ans...)
	}
	if promotion > 0 {
		ans = append([]byte{promotion + '0'}, ans...)
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
