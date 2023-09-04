package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when859_thenSuccess(t *testing.T) {
	assert.Equal(t, true, buddyStrings("ab", "ba"))
	assert.Equal(t, false, buddyStrings("ab", "ab"))
	assert.Equal(t, true, buddyStrings("aa", "aa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func buddyStrings(s string, goal string) bool {
	m, n := len(s), len(goal)
	if m != n {
		return false
	}
	swapped := 0
	l, r := 0, n-1
	charMap := [26]byte{}
	duplicate := false
	for l <= r {
		if s[l] == goal[l] {
			charMap[s[l]-'a'] += 1
			if charMap[s[l]-'a'] > 1 {
				duplicate = true
			}
			l += 1
			continue
		}
		if s[r] == goal[r] {
			r -= 1
			continue
		}
		if s[l] == goal[r] && s[r] == goal[l] {
			l, r = l+1, r-1
			swapped += 1
		} else {
			break
		}
	}
	if l <= r || swapped > 1 {
		return false
	}
	if swapped == 0 {
		return duplicate
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
