package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when744_thenSuccess(t *testing.T) {
	assert.Equal(t, byte('c'), nextGreatestLetter([]byte("cfg"), 'a'))
	assert.Equal(t, byte('f'), nextGreatestLetter([]byte("cfg"), 'c'))
	assert.Equal(t, byte('c'), nextGreatestLetter([]byte("cfg"), 'k'))
}

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreatestLetter(letters []byte, target byte) byte {
	// 2 分
	ans := letters[0]
	for l, r := 0, len(letters)-1; l <= r; {
		mid := l + (r-l)>>1
		if letters[mid] > target {
			ans = letters[mid]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
