package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when557_thenSuccess(t *testing.T) {
	ans := reverseWords("Let's take LeetCode contest")
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	bytes := []byte(s)
	length := len(bytes)
	l, r := 0, 0
	for i := 0; i < length; i++ {
		space := bytes[i] == ' '
		if !space {
			// 非空格即移动右指针
			r = i
			if i == 0 || bytes[i-1] == ' ' {
				// 非空格，且前一个字符空格，移动左指针
				l = i
			}
		}
		// 遇到空格或字符串结束，反转 [l,r]
		if space || i == length-1 {
			for l < r {
				bytes[l], bytes[r] = bytes[r], bytes[l]
				l, r = l+1, r-1
			}
		}
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
