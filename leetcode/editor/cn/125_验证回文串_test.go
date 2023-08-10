package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when125_thenSuccess(t *testing.T) {
	palindrome := isPalindrome125("A man, a plan, a canal: Panama")
	assert.Equal(t, true, palindrome)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome125(s string) bool {
	ans := true
	left, right := 0, len(s)-1
	for left < right {
		if !isAlNum(s[left]) {
			left += 1
			continue
		}
		if !isAlNum(s[right]) {
			right -= 1
			continue
		}
		if lower(s[left]) != lower(s[right]) {
			ans = false
			break
		}
		left, right = left+1, right-1
	}
	return ans
}

func isAlNum(u uint8) bool {
	if u >= '0' && u <= '9' {
		return true
	}
	if u >= 'A' && u <= 'Z' {
		return true
	}
	if u >= 'a' && u <= 'z' {
		return true
	}
	return false
}

func lower(u uint8) uint8 {
	if u >= 'A' && u <= 'Z' {
		return u + ('a' - 'A')
	}
	return u
}

//leetcode submit region end(Prohibit modification and deletion)
